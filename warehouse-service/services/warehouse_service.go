package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/putrapratamanst/ecommerce/warehouse-service/messaging"
	"github.com/putrapratamanst/ecommerce/warehouse-service/models"
	"github.com/putrapratamanst/ecommerce/warehouse-service/repositories"
	"github.com/putrapratamanst/ecommerce/warehouse-service/utils"
)

type WarehouseService interface {
	CreateWarehouse(warehouse *models.Warehouse) error
	GetWarehouseByID(id uint) (*models.Warehouse, error)
	SetWarehouseShop(warehouseID string, shopID string) error
	GetShopsWarehouse(warehouseID string) (*[]models.WarehouseShop, error)
	GetWarehousesByShopID(shopID uint) ([]models.Warehouse, error)
	SetWarehouseStatus(warehouseID uint, isActive bool) error
    AdjustStockForOrder(orderID uint) error
    TransferStock(fromWarehouseID, toWarehouseID, productID uint, quantity int) error
    AddStock(warehouseID, productID uint, quantity int) error
    StartListening(rabbitMQ *messaging.RabbitMQ)
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

func (s *warehouseService) GetWarehouseByID(id uint) (*models.Warehouse, error) {
	return s.warehouseRepository.GetWarehouseByID(id)
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

func (s *warehouseService) GetShopsWarehouse(warehouseID string) (*[]models.WarehouseShop, error) {
	warehouseIDInt, _ := strconv.ParseUint(warehouseID, 10, 32)

	getShop, err := s.warehouseRepository.GetShopByWarehouseID(uint(warehouseIDInt))
	if err != nil {
		return nil, err
	}

	return getShop, nil

}

func (w *warehouseService) GetWarehousesByShopID(shopID uint) ([]models.Warehouse, error) {
    return w.warehouseRepository.GetWarehousesByShopID(shopID)
}

func (w *warehouseService) StartListening(rabbitMQ *messaging.RabbitMQ) {
    msgChan, err := rabbitMQ.Channel.Consume(
        "order_placed", // queue name
        "",             // consumer name
        true,           // auto-ack
        false,          // exclusive
        false,          // no-local
        false,          // no-wait
        nil,            // arguments
    )

    if err != nil {
        log.Fatalf("Failed to register a consumer: %v", err)
    }

    for msg := range msgChan {
        log.Printf("Received a message: %s", msg.Body)

        orderID, err := strconv.ParseUint(string(msg.Body), 10, 32)
        if err != nil {
            log.Printf("Failed to parse order ID: %v", err)
            continue
        }

        err = w.AdjustStockForOrder(uint(orderID))
        if err != nil {
            log.Printf("Failed to adjust stock for order %d: %v", orderID, err)
            continue
        }
    }
}

func (w *warehouseService) AdjustStockForOrder(orderID uint) error {
    orderDetails, err := w.fetchOrderDetails(orderID)
    if err != nil {
        return err
    }

    for _, item := range orderDetails.Items {
        err := w.warehouseRepository.UpdateStock(item.WarehouseID, item.ProductID, -item.Quantity)
        if err != nil {
            log.Printf("Failed to update stock for ProductID %d in WarehouseID %d: %v", item.ProductID, item.WarehouseID, err)
            return err
        }
    }

    return nil
}

func (w *warehouseService) AddStock(warehouseID, productID uint, quantity int) error {
    warehouse, err := w.GetWarehouseByID(warehouseID)
    if err != nil || !warehouse.IsActive {
        return errors.New("warehouse is inactive or not found")
    }
    return w.warehouseRepository.UpdateStock(warehouseID, productID, quantity)
}

func (w *warehouseService) SetWarehouseStatus(warehouseID uint, isActive bool) error {
    warehouse, err := w.warehouseRepository.GetWarehouseByID(warehouseID)
    if err != nil {
        return err
    }
    warehouse.IsActive = isActive
    return w.warehouseRepository.SetStatus(warehouse)
}

func (w *warehouseService) TransferStock(fromWarehouseID, toWarehouseID, productID uint, quantity int) error {
    fromWarehouse, err := w.warehouseRepository.GetWarehouseByID(fromWarehouseID)
    if err != nil || !fromWarehouse.IsActive {
        return errors.New("source warehouse is inactive or not found")
    }

    toWarehouse, err := w.warehouseRepository.GetWarehouseByID(toWarehouseID)
    if err != nil || !toWarehouse.IsActive {
        return errors.New("destination warehouse is inactive or not found")
    }

    return w.warehouseRepository.TransferStock(fromWarehouseID, toWarehouseID, productID, quantity)
}

// Fetch order details from Order Service
func (w *warehouseService) fetchOrderDetails(orderID uint) (*utils.OrderDetails, error) {
    // Replace with your actual Order Service URL
    orderServiceURL := fmt.Sprintf("http://order-service/orders/%d", orderID)

    resp, err := http.Get(orderServiceURL)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, errors.New("failed to fetch order details")
    }

    var orderDetails utils.OrderDetails
    if err := json.NewDecoder(resp.Body).Decode(&orderDetails); err != nil {
        return nil, err
    }

    return &orderDetails, nil
}