package db

import (
	"example/web-service-gin/models"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error

	dbDriver := os.Getenv("DB_DRIVER")
	dbConnection := os.Getenv("DB_CONNECTION")

	switch dbDriver {
	case "mysql":
		DB, err = gorm.Open(mysql.Open(dbConnection), &gorm.Config{})
	default:
		log.Fatalf("Unsupported DB_DRIVER: %v", dbDriver)
	}

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto migrate the models
	DB.AutoMigrate(&models.Product{})
}
