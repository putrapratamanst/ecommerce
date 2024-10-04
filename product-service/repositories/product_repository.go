package repositories

import (
	"github.com/putrapratamanst/ecommerce/product-service/models"
	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

func (r *ProductRepository) FindAllProducts(limit, offset int) ([]models.Product, int64, error) {
	var products []models.Product
	var total int64
	if err := r.DB.Limit(limit).Offset(offset).Find(&products).Error; err != nil {
		return nil, 0, err
	}
	r.DB.Model(&models.Product{}).Count(&total)
	return products, total, nil
}
