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

func (r *WarehouseRepository) FindByID(id string) (*models.Warehouse, error) {
	var warehouse models.Warehouse
	err := r.db.First(&warehouse, id).Error
	return &warehouse, err
}

func (r *WarehouseRepository) SetWarehouseShop(warehouseShop *models.WarehouseShop) error {
	return r.db.Create(warehouseShop).Error
}

func (r *WarehouseRepository) FindByShopIDAndWarehouseID(warehouseID uint, shopID uint) (*models.WarehouseShop, error) {
	var warehouseShop models.WarehouseShop
	err := r.db.First(&warehouseShop, "warehouse_id = ? AND shop_id = ?", warehouseID, shopID).Error
	return &warehouseShop, err
}
