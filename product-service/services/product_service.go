package services

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/putrapratamanst/ecommerce/product-service/models"
	"github.com/putrapratamanst/ecommerce/product-service/repositories"
)

var ctx = context.Background()

type ProductService struct {
	ProductRepository *repositories.ProductRepository
	RedisClient       *redis.Client
}

func NewProductService(repo *repositories.ProductRepository, redisClient *redis.Client) *ProductService {
	return &ProductService{
		ProductRepository: repo,
		RedisClient:       redisClient,
	}
}

func (s *ProductService) GetProducts(limit, offset int) ([]models.Product, int64, error) {
	cacheKey := fmt.Sprintf("products_%d_%d", limit, offset)

	// Try to get products from Redis cache
	cachedProducts, err := s.RedisClient.Get(ctx, cacheKey).Result()
	if err == redis.Nil {
		// If not in cache, fetch from database
		products, total, err := s.ProductRepository.FindAllProducts(limit, offset)
		if err != nil {
			return nil, 0, err
		}

		// Store the fetched products in Redis with an expiration time
		productsJSON, _ := json.Marshal(products)
		s.RedisClient.Set(ctx, cacheKey, productsJSON, 10*time.Minute)

		return products, total, nil
	} else if err != nil {
		return nil, 0, err
	}

	// If found in cache, decode the cached data
	var products []models.Product
	if err := json.Unmarshal([]byte(cachedProducts), &products); err != nil {
		return nil, 0, err
	}

	// Since total isn't cached, fetch total from database
	var total int64
	s.ProductRepository.DB.Model(&models.Product{}).Count(&total)

	return products, total, nil
}
