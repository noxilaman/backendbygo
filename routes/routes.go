package routes

import (
	"example/web-service-gin/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Product routes
	r.GET("/products", controllers.GetProducts)
	r.POST("/products", controllers.CreateProduct)

	return r
}
