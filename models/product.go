package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name  string  `json:"name" binding:"required,min=3,max=100" gorm:"unique"`
	Price float64 `json:"price"`
}
