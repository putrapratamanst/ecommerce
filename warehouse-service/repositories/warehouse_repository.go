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

func (r *WarehouseRepository) ReleaseStock(warehouseID int, productID int, quantity int) error {
    return r.db.Transaction(func(tx *gorm.DB) error {
        var stock models.WarehouseProductStock
        if err := tx.Model(&stock).Where("warehouse_id = ? AND product_id = ?", warehouseID, productID).First(&stock).Error; err != nil {
            return err
        }
        stock.Quantity += quantity
        return tx.Save(&stock).Error
    })
}

func (r *WarehouseRepository) TransferStock(fromWarehouseID int, toWarehouseID int, productID int, quantity int) error {
    return r.db.Transaction(func(tx *gorm.DB) error {
        var fromStock models.WarehouseProductStock
        if err := tx.Model(&fromStock).Where("warehouse_id = ? AND product_id = ? AND quantity >= ?", fromWarehouseID, productID, quantity).First(&fromStock).Error; err != nil {
            return err // Insufficient stock or not found
        }

        fromStock.Quantity -= quantity
        if err := tx.Save(&fromStock).Error; err != nil {
            return err
        }

        var toStock models.WarehouseProductStock
        if err := tx.Model(&toStock).Where("warehouse_id = ? AND product_id = ?", toWarehouseID, productID).First(&toStock).Error; err != nil {
            toStock = models.WarehouseProductStock{
                WarehouseID: toWarehouseID,
                ProductID:   productID,
                Quantity:    quantity,
            }
            return tx.Create(&toStock).Error
        }

        toStock.Quantity += quantity
        return tx.Save(&toStock).Error
    })
}

func (r *WarehouseRepository) SetWarehouseActive(warehouseID int, active bool) error {
    return r.db.Model(&models.Warehouse{}).Where("id = ?", warehouseID).Update("active", active).Error
}

func (r *WarehouseRepository) GetStock(warehouseID int, productID int) (int, error) {
    var stock models.WarehouseProductStock
    if err := r.db.Where("warehouse_id = ? AND product_id = ?", warehouseID, productID).First(&stock).Error; err != nil {
        return 0, err
    }
    return stock.Quantity, nil
}

func (r *WarehouseRepository) Create(warehouse *models.Warehouse) error {
	return r.db.Create(warehouse).Error
}

func (r *WarehouseRepository) FindByID(id string) (*models.Warehouse, error) {
	var warehouse models.Warehouse
	err := r.db.Where("status = ?", "active").First(&warehouse, id).Error
	return &warehouse, err
}

