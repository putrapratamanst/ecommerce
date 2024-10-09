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
	ProductID   int
	WarehouseID int
	Quantity    int
}

type WarehouseShop struct {
	gorm.Model
	ShopID      uint
	WarehouseID uint `gorm:"notNull;size:40" json:"-"`

	Warehouse *Warehouse `gorm:"foreignKey:WarehouseID" json:"Warehouse,omitempty"`
}
