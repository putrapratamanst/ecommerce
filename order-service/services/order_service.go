package services

import (
	"context"
	"encoding/json"
	"time"

	"github.com/putrapratamanst/ecommerce/order-service/messaging"
	"github.com/putrapratamanst/ecommerce/order-service/models"
	"github.com/putrapratamanst/ecommerce/order-service/repositories"
)
type OrderService struct {
    repo *repositories.OrderRepository
    mq   *messaging.RabbitMQ
}

func NewOrderService(repo *repositories.OrderRepository, mq *messaging.RabbitMQ) *OrderService {
    return &OrderService{repo: repo, mq: mq}
}

func (s *OrderService) CheckoutOrder(ctx context.Context, order *models.Order) error {
    err := s.repo.CreateOrder(order)
    if err != nil {
        return err
    }

    reservationMessage := map[string]interface{}{
        "order_id":     order.ID,
        "shop_id":      order.ShopID,
        "warehouse_id": order.WarehouseID,
        "quantity":     order.Quantity,
    }

    messageBytes, _ := json.Marshal(reservationMessage)
    err = s.mq.Publish("stock_reservation_queue", messageBytes)
    if err != nil {
        return err
    }

    // Set timer to release stock if not paid
    go s.releaseStockIfNotPaid(ctx, order)

    return nil
}

func (s *OrderService) releaseStockIfNotPaid(ctx context.Context, order *models.Order) {
    select {
    case <-ctx.Done():
        return
    case <-time.After(15 * time.Minute):
        releaseMessage := map[string]interface{}{
            "order_id":   order.ID,
            "product_id": order.ProductID,
            "quantity":   order.Quantity,
        }
        messageBytes, _ := json.Marshal(releaseMessage)
        s.mq.Publish("stock_release_queue", messageBytes)
    }
}

func (s *OrderService) UpdateOrderStatus(order *models.Order) error {
    return s.repo.UpdateOrderStatus(order.ID, order.Status)
}

func (s *OrderService) GetOrderByID(orderID int) (*models.Order, error) {
    return s.repo.FindOrderByID(orderID)
}

func (s *OrderService) ReleaseOrder(ctx context.Context, order *models.Order) error {
    releaseMessage := map[string]interface{}{
        "order_id":     order.ID,
        "warehouse_id": order.WarehouseID,
        "product_id":   order.ProductID,
        "quantity":     order.Quantity,
    }

    messageBytes, _ := json.Marshal(releaseMessage)
    err := s.mq.Publish("stock_release_queue", messageBytes)
    if err != nil {
        return err
    }

    return nil
}
