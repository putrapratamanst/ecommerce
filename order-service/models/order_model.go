package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID      uint    `gorm:"not null"`
	ProductID   uint    `gorm:"not null" validate:"required"`
	Quantity    int     `gorm:"not null" validate:"required"`
	TotalPrice  float64 `gorm:"type:decimal(10,2)"`
	Status      string  `gorm:"type:varchar(20);default:'active'"`
	ShopID      int     `gorm:"not null" validate:"required"`
	WarehouseID int     `gorm:"not null" validate:"required"`
}
