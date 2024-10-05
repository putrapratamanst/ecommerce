package models

import "gorm.io/gorm"

type Warehouse struct {
	gorm.Model
	Name     string `json:"name" gorm:"not null" validate:"required"`
	Location string `json:"location" validate:"required"`
	Status   string `json:"status" validate:"required"`
}

type WarehouseShop struct {
	gorm.Model
	ShopID      uint
	WarehouseID uint `gorm:"notNull;size:40" json:"-"`

	Warehouse *Warehouse `gorm:"foreignKey:WarehouseID" json:"Warehouse,omitempty"`
}

type WarehouseShopProductStock struct {
	gorm.Model
	ShopID        uint
	ProductID     uint
	StockQuantity int  `json:"-"`
	WarehouseID   uint `gorm:"notNull;size:40" json:"-"`

	Warehouse *Warehouse `gorm:"foreignKey:WarehouseID" json:"Warehouse,omitempty"`
}
