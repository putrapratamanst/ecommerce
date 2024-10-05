package services

import (
	"strconv"

	"github.com/putrapratamanst/ecommerce/warehouse-service/models"
	"github.com/putrapratamanst/ecommerce/warehouse-service/repositories"
)

type WarehouseService interface {
	CreateWarehouse(warehouse *models.Warehouse) error
	GetWarehouseByID(id string) (*models.Warehouse, error)
	SetWarehouseShop(warehouseID string, shopID string) error
	GetWarehouseShop(warehouseID string, shopID string) (*models.WarehouseShop, error)
}

type warehouseService struct {
	warehouseRepository *repositories.WarehouseRepository
}

func NewWarehouseService(warehouseRepository *repositories.WarehouseRepository) WarehouseService {
	return &warehouseService{
		warehouseRepository: warehouseRepository,
	}
}

func (s *warehouseService) CreateWarehouse(warehouse *models.Warehouse) error {
	return s.warehouseRepository.Create(warehouse)
}

func (s *warehouseService) GetWarehouseByID(id string) (*models.Warehouse, error) {
	return s.warehouseRepository.FindByID(id)
}

func (s *warehouseService) GetWarehouseShop(warehouseID string, shopID string) (*models.WarehouseShop, error) {
	warehouseIDInt, _ := strconv.ParseUint(warehouseID, 10, 32)
	shopIDInt, _ := strconv.ParseUint(shopID, 10, 32)	

	return s.warehouseRepository.FindByShopIDAndWarehouseID(uint(warehouseIDInt), uint(shopIDInt))
}

func (s *warehouseService) SetWarehouseShop(warehouseID string, shopID string) error {
	warehouseIDInt, _ := strconv.ParseUint(warehouseID, 10, 32)
	shopIDInt, _ := strconv.ParseUint(shopID, 10, 32)

	warehouseShop := models.WarehouseShop{
		WarehouseID: uint(warehouseIDInt),
		ShopID:      uint(shopIDInt),
	}
	return s.warehouseRepository.SetWarehouseShop(&warehouseShop)
}
