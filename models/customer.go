package models

import (
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Name         string `json:"name" binding:"required,min=3,max=100" gorm:"unique"`
	Desc         string `json:"desc"`
	ContactName  string `json:"contact_name"`
	ContactPhone string `json:"contact_phone"`
	ContactEmail string `json:"contact_email"`
}
