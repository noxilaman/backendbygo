package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" binding:"required,min=3,max=100" gorm:"unique"`
	Password string `json:"password" binding:"required,min=3,max=100"`
	Name     string `json:"name" binding:"required,min=3,max=100"`
	Role     string `json:"role" binding:"required,min=3,max=100"`
	Token    string `json:"token"`
	Retoken  string `json:"retoken"`
}
