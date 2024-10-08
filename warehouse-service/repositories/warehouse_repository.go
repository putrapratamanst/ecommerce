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

func (r *WarehouseRepository) SetWarehouseShop(warehouseShop *models.WarehouseShop) error {
	return r.db.Create(warehouseShop).Error
}

func (r *WarehouseRepository) GetWarehouseByID(id uint) (*models.Warehouse, error) {
    var warehouse models.Warehouse
    result := r.db.First(&warehouse, id)
    return &warehouse, result.Error
}

func (r *WarehouseRepository) GetWarehousesByShopID(shopID uint) ([]models.Warehouse, error) {
    var warehouses []models.Warehouse

    var warehouseShopEntries []models.WarehouseShop
    result := r.db.Where("shop_id = ?", shopID).Find(&warehouseShopEntries)
    if result.Error != nil {
        return nil, result.Error
    }

    for _, entry := range warehouseShopEntries {
        var warehouse models.Warehouse
        if err := r.db.First(&warehouse, entry.WarehouseID).Error; err == nil {
            warehouses = append(warehouses, warehouse)
        }
    }

    return warehouses, nil
}

func (r *WarehouseRepository) GetProductStock(warehouseID, productID uint) (*models.WarehouseShopProductStock, error) {
    var stock models.WarehouseShopProductStock
    result := r.db.Where("warehouse_id = ? AND product_id = ?", warehouseID, productID).First(&stock)
    return &stock, result.Error
}

func (r *WarehouseRepository) UpdateStock(warehouseID, productID uint, quantity int) error {
    stock, err := r.GetProductStock(warehouseID, productID)
    if err != nil {
        return err
    }
    stock.StockQuantity += quantity
    return r.db.Save(&stock).Error
}

func (r *WarehouseRepository) TransferStock(fromWarehouseID, toWarehouseID, productID uint, quantity int) error {
    err := r.UpdateStock(fromWarehouseID, productID, -quantity)
    if err != nil {
        return err
    }
    return r.UpdateStock(toWarehouseID, productID, quantity)
}

func (r *WarehouseRepository) GetByShopIDAndWarehouseID(warehouseID uint, shopID uint) (*models.WarehouseShop, error) {
	var warehouseShop models.WarehouseShop
	err := r.db.First(&warehouseShop, "warehouse_id = ? AND shop_id = ? AND status = active", warehouseID, shopID).Error
	return &warehouseShop, err
}

func (r *WarehouseRepository) GetShopByWarehouseID(warehouseID uint) (*[]models.WarehouseShop, error) {
	var warehouseShop []models.WarehouseShop
	err := r.db.Find(&warehouseShop, "warehouse_id = ?", warehouseID).Error
	return &warehouseShop, err
}

func (r *WarehouseRepository) SetStatus(warehouse *models.Warehouse) error {
    return r.db.Save(&warehouse).Error
}

