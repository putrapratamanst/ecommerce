package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID            uint      `gorm:"not null"`
	ProductID         uint      `gorm:"not null"`
	Quantity          int       `gorm:"not null"`
	TotalPrice        float64   `gorm:"type:decimal(10,2)"`
	Status            string    `gorm:"type:varchar(20);default:'pending'"`
	PaymentExpiration time.Time `gorm:"not null"`
}
