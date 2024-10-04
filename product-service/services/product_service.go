package services

import (
	"github.com/putrapratamanst/ecommerce/product-service/models"
	"github.com/putrapratamanst/ecommerce/product-service/repositories"
)

type ProductService struct {
    ProductRepository *repositories.ProductRepository
}

func NewProductService(repo *repositories.ProductRepository) *ProductService {
    return &ProductService{ProductRepository: repo}
}

func (s *ProductService) GetProducts(limit, offset int) ([]models.Product, int64, error) {
    return s.ProductRepository.FindAllProducts(limit, offset)
}
