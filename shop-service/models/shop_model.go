package models

import "gorm.io/gorm"

type Shop struct {
	gorm.Model
	Name    string `json:"name" gorm:"not null" validate:"required"`
	OwnerID int    `json:"owner_id"`
}
