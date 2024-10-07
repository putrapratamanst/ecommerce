package repositories

import (
	"time"

	"github.com/putrapratamanst/ecommerce/order-service/models"
	"gorm.io/gorm"
)

type OrderRepository interface {
	CreateOrder(order *models.Order) error
	FindOrderByID(id uint) (*models.Order, error)
	FindOrdersByUserID(userID uint) ([]models.Order, error)
	UpdateOrderStatus(orderID uint, status string) error
	FindExpiredOrders() ([]models.Order, error)
}

type orderRepositoryImpl struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepositoryImpl{db}
}

func (r *orderRepositoryImpl) CreateOrder(order *models.Order) error {
	return r.db.Create(order).Error
}

func (r *orderRepositoryImpl) FindOrderByID(id uint) (*models.Order, error) {
	var order models.Order
	result := r.db.First(&order, id)
	return &order, result.Error
}

func (r *orderRepositoryImpl) FindOrdersByUserID(userID uint) ([]models.Order, error) {
	var orders []models.Order
	result := r.db.Where("user_id = ?", userID).Find(&orders)
	return orders, result.Error
}

func (r *orderRepositoryImpl) UpdateOrderStatus(orderID uint, status string) error {
	return r.db.Model(&models.Order{}).Where("id = ?", orderID).Update("status", status).Error
}

func (r *orderRepositoryImpl) FindExpiredOrders() ([]models.Order, error) {
    var expiredOrders []models.Order
    now := time.Now()
    result := r.db.Where("status = ? AND payment_expiration < ?", "pending", now).Find(&expiredOrders)
    return expiredOrders, result.Error
}