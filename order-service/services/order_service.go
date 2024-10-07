package services

import (
	"github.com/putrapratamanst/ecommerce/order-service/messaging"
	"github.com/putrapratamanst/ecommerce/order-service/models"
	"github.com/putrapratamanst/ecommerce/order-service/repositories"
)

type OrderService interface {
    PlaceOrder(userID uint, productID uint, quantity int) (*models.Order, error)
	GetOrderByID(orderID uint) (*models.Order, error)
    GetOrdersByUserID(userID uint) ([]models.Order, error)
}

type orderService struct {
    orderRepo repositories.OrderRepository
    rabbitMQ  *messaging.RabbitMQ // RabbitMQ instance for publishing messages
}

func NewOrderService(orderRepo repositories.OrderRepository, rabbitMQ *messaging.RabbitMQ) OrderService {
    return &orderService{
        orderRepo: orderRepo,
        rabbitMQ:  rabbitMQ,
    }
}

// PlaceOrder creates an order and publishes an OrderPlaced event
func (s *orderService) PlaceOrder(userID uint, productID uint, quantity int) (*models.Order, error) {
    order := &models.Order{
        UserID:     userID,
        ProductID:  productID,
        Quantity:   quantity,
        TotalPrice: float64(quantity) * 10.00,
        Status:     "pending",
    }

    err := s.orderRepo.CreateOrder(order)
    if err != nil {
        return nil, err
    }

    err = s.rabbitMQ.PublishOrderPlaced(order.ID)
    if err != nil {
        return nil, err
    }

    return order, nil
}

func (s *orderService) GetOrderByID(orderID uint) (*models.Order, error) {
    return s.orderRepo.FindOrderByID(orderID)
}

func (s *orderService) GetOrdersByUserID(userID uint) ([]models.Order, error) {
    return s.orderRepo.FindOrdersByUserID(userID)
}
