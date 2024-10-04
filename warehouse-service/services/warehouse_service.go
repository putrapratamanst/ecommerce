package services

import (
	"github.com/putrapratamanst/ecommerce/warehouse-service/models"
	"github.com/putrapratamanst/ecommerce/warehouse-service/repositories"
)

type WarehouseService interface {
	CreateWarehouse(warehouse *models.Warehouse) error
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