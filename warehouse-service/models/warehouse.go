package models

import (
	"gorm.io/gorm"
)

type Warehouse struct {
	gorm.Model
    Name      string    `gorm:"not null"`
    Location  string    `gorm:"not null"`
    IsActive  bool      `gorm:"default:true"`
}
type WarehouseStock struct {
	gorm.Model
    WarehouseID uint      `gorm:"not null"`
    ProductID   uint      `gorm:"not null"`
    Quantity    int       `gorm:"not null"`
}

type StockTransfer struct {
	gorm.Model
    FromWarehouseID   uint      `gorm:"not null"`
    ToWarehouseID     uint      `gorm:"not null"`
    ProductID         uint      `gorm:"not null"`
    Quantity          int       `gorm:"not null"`
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
