package routes

import (
	"example/web-service-gin/controllers"
	"example/web-service-gin/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/login", controllers.Login)
	r.POST("/register", controllers.RegisterUser)
	r.POST("/refresh", controllers.RefreshToken)
	// Product routes
	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/users/:id", controllers.GetUser)

		protected.GET("/products", controllers.GetProducts)
		protected.POST("/products", controllers.CreateProduct)
		protected.GET("/products/:id", controllers.GetProduct)
		protected.PUT("/products/:id", controllers.UpdateProduct)
		protected.DELETE("/products/:id", controllers.DeleteProduct)

		protected.GET("/customers", controllers.GetCustomers)
		protected.POST("/customers", controllers.CreateCustomer)
		protected.GET("/customers/:id", controllers.GetCustomer)
		protected.PUT("/customers/:id", controllers.UpdateCustomer)
		protected.DELETE("/customers/:id", controllers.DeleteCustomer)
	}

	return r
}
