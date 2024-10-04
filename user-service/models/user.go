package models

import (
	"gorm.io/gorm"
)

type User struct {
    gorm.Model
    Name     string `json:"name"`
    Email    string `json:"email" gorm:"unique" validate:"required,email"`
    Phone    string `json:"phone" gorm:"unique" validate:"required,numeric,max=12"`
    Password string `json:"password" validate:"required,min=6"`
}
