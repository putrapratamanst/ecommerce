package services

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/putrapratamanst/ecommerce/warehouse-service/messaging"
	"github.com/putrapratamanst/ecommerce/warehouse-service/models"
	"github.com/putrapratamanst/ecommerce/warehouse-service/repositories"
)

type WarehouseService struct {
	repo *repositories.WarehouseRepository
}

func NewWarehouseService(repo *repositories.WarehouseRepository) *WarehouseService {
	return &WarehouseService{repo: repo}
}

func (s *WarehouseService) HandleStockReservation(warehouseID int, productID int, quantity int) error {
	return s.repo.AdjustStock(warehouseID, productID, -quantity)
}

func (s *WarehouseService) ListenForStockReservation(mq *messaging.RabbitMQ) {
	mq.Consume("stock_reservation_queue", func(message []byte) {
		var reservationData map[string]interface{}
		json.Unmarshal(message, &reservationData)

		warehouseID := int(reservationData["warehouse_id"].(float64))
		productID := int(reservationData["product_id"].(float64))
		quantity := int(reservationData["quantity"].(float64))

		err := s.HandleStockReservation(warehouseID, productID, quantity)
		if err != nil {
			fmt.Println(err.Error())
		}
	})
}

func (s *WarehouseService) ListenForStockRelease(mq *messaging.RabbitMQ) {
	mq.Consume("stock_release_queue", func(message []byte) {
		var releaseData map[string]interface{}
		json.Unmarshal(message, &releaseData)

		warehouseID := int(releaseData["warehouse_id"].(float64))
		productID := int(releaseData["product_id"].(float64))
		quantity := int(releaseData["quantity"].(float64))

		err := s.HandleStockRelease(warehouseID, productID, quantity)
		if err != nil {
			fmt.Println(err.Error())
		}
	})
}

func (s *WarehouseService) HandleStockRelease(warehouseID int, productID int, quantity int) error {
	return s.repo.ReleaseStock(warehouseID, productID, quantity) // Merilis stok dari gudang yang sesuai
}

func (s *WarehouseService) CreateWarehouse(warehouse *models.Warehouse) error {
	return s.repo.Create(warehouse)
}

func (s *WarehouseService) GetWarehouseByID(id uint) (*models.Warehouse, error) {
	return s.repo.FindByID(id)
}

func (s *WarehouseService) SetWarehouseShop(warehouseID string, shopID string) error {
	shopIDInt, _ := strconv.Atoi(shopID)
	warehouseIDInt, _ := strconv.Atoi(warehouseID)

    // skip if warehouseID or shopID is exist
    getWarehouseShop, _ := s.repo.GetWarehouseShop(uint(warehouseIDInt), uint(shopIDInt))
    if getWarehouseShop != nil {
        return nil
    }
    fmt.Println("masuk sini harusnya")
	return s.repo.SetWarehouseShop(uint(warehouseIDInt), uint(shopIDInt))
}
func (s *WarehouseService) SetWarehouseActive(warehouseID int, active bool) error {
	return s.repo.SetWarehouseActive(warehouseID, active)
}

func (s *WarehouseService) AdjustStock(warehouseID, productID, quantity int) error {
	return s.repo.AdjustStock(warehouseID, productID, quantity)
}
func (s *WarehouseService) TransferStock(fromWarehouseID, toWarehouseID, productID, quantity int) error {
	return s.repo.TransferStock(fromWarehouseID, toWarehouseID, productID, quantity)
}
