package main

import (
	"example/web-service-gin/db"
	"example/web-service-gin/routes"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}

	db.InitDB()

	// Initialize the router
	r := routes.SetupRouter()

	// Get the port from environment variables
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}

	// Run the server
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
