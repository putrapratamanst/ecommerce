package repositories

import (
	"log"

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

func (r *WarehouseRepository) FindByID(id uint) (*models.Warehouse, error) {
	var warehouse models.Warehouse
	err := r.db.Where("active = ?", "true").First(&warehouse, id).Error
	return &warehouse, err
}

func (r *WarehouseRepository) SetWarehouseShop(warehouseID uint, shopID uint) error {
	var warehouseShop models.WarehouseShop
	warehouseShop.WarehouseID = warehouseID
	warehouseShop.ShopID = shopID
	return r.db.Create(&warehouseShop).Error
}

func (r *WarehouseRepository) AdjustStock(warehouseID int, productID int, quantity int) error {
    tx := r.db.Begin()
    if tx.Error != nil {
        return tx.Error
    }

    defer func() {
        if r := recover(); r != nil {
            tx.Rollback()
            log.Println("Transaction rollback due to panic:", r)
        }
    }()

    // First, lock the row with SELECT FOR UPDATE to avoid race conditions
    if err := tx.Exec(`
        SELECT quantity 
        FROM warehouse_product_stocks 
        WHERE warehouse_id = ? AND product_id = ? 
        FOR UPDATE`, warehouseID, productID).Error; err != nil {
        tx.Rollback()
        log.Println("Error locking row:", err)
        return err
    }

    // Update the stock quantity
    result := tx.Exec(`
        UPDATE warehouse_product_stocks
        SET quantity = quantity + ? 
        WHERE warehouse_id = ? AND product_id = ?`, quantity, warehouseID, productID)
    
    if result.Error != nil {
        tx.Rollback()
        log.Println("Error updating stock:", result.Error)
        return result.Error
    }

    // If no rows were updated, insert a new row
    if result.RowsAffected == 0 {
        log.Println("No rows updated, inserting new record.")
        if err := tx.Exec(`
            INSERT INTO warehouse_product_stocks (warehouse_id, product_id, quantity) 
            VALUES (?, ?, ?)`, warehouseID, productID, quantity).Error; err != nil {
            tx.Rollback()
            log.Println("Error inserting new stock record:", err)
            return err
        }
    }

    // Commit the transaction
    if err := tx.Commit().Error; err != nil {
        log.Println("Transaction commit failed:", err)
        return err
    }

    log.Println("Stock successfully updated for warehouse:", warehouseID, "product:", productID)
    return nil
}

func (r *WarehouseRepository) GetWarehouseShop(warehouseID, shopID uint) (*models.WarehouseShop, error) {
	var ws models.WarehouseShop
	if err := r.db.Where("warehouse_id = ? AND shop_id = ?", warehouseID, shopID).First(&ws).Error; err != nil {
		return nil, err
	}
	return &ws, nil
}
