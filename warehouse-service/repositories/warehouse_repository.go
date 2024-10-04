package repositories

import (
	"github.com/putrapratamanst/ecommerce/warehouse-service/models"
	"gorm.io/gorm"
)

type WarehouseRepository struct {
	db *gorm.DB
}

func NewWarehouseRepository(db *gorm.DB) *WarehouseRepository {
	return &WarehouseRepository{db: db}
}

func (r *WarehouseRepository) Create(warehouse *models.Warehouse) error {
	return r.db.Create(warehouse).Error
}