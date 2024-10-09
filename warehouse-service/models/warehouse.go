package models

import (
	"gorm.io/gorm"
)

type Warehouse struct {
	gorm.Model
	ShopID   int // Menyimpan ID toko
	Name     string
	Location string
	Active   bool
}

type WarehouseProductStock struct {
	gorm.Model
	ProductID   int `gorm:"primaryKey"`
	WarehouseID int `gorm:"primaryKey"`
	Quantity    int
}
