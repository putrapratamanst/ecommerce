package repositories

import (
	"github.com/putrapratamanst/ecommerce/order-service/models"
	"gorm.io/gorm"
)

type OrderRepository struct {
    db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
    return &OrderRepository{db: db}
}

func (r *OrderRepository) CreateOrder(order *models.Order) error {
    return r.db.Create(order).Error
}

func (r *OrderRepository) UpdateOrderStatus(orderID uint, status string) error {
    return r.db.Model(&models.Order{}).Where("id = ?", orderID).Update("status", status).Error
}

func (r *OrderRepository) FindOrderByID(orderID int) (*models.Order, error) {
    var order models.Order
    if err := r.db.First(&order, orderID).Error; err != nil {
        return &order, err
    }
    return &order, nil
}

