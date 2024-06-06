package controllers

import (
	"example/web-service-gin/db"
	"example/web-service-gin/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetCustomers(c *gin.Context) {
	var customers []models.Customer

	// Get page and pageSize query parameters
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
		return
	}

	pageSize, err := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	if err != nil || pageSize < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page size"})
		return
	}

	// Calculate offset
	offset := (page - 1) * pageSize

	// Get total count of customers
	var totalCount int64
	if err := db.DB.Model(&models.Customer{}).Count(&totalCount).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch customers count"})
		return
	}

	// Fetch products with limit and offset
	if err := db.DB.Limit(pageSize).Offset(offset).Find(&customers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch customers"})
		return
	}

	// Calculate total pages
	totalPages := (int(totalCount) + pageSize - 1) / pageSize

	// Prepare response
	response := gin.H{
		"data":       customers,
		"page":       page,
		"pageSize":   pageSize,
		"total":      totalCount,
		"totalPages": totalPages,
	}

	c.JSON(http.StatusOK, response)
}

func GetCustomer(c *gin.Context) {
	var customer models.Customer
	if err := db.DB.First(&customer, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "customer not found"})
		return
	}
	c.JSON(http.StatusOK, customer)
}

func CreateCustomer(c *gin.Context) {
	var customer models.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Create(&customer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create customer"})
		return
	}

	c.JSON(http.StatusCreated, customer)
}

func UpdateCustomer(c *gin.Context) {
	var customer models.Customer
	if err := db.DB.First(&customer, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Save(&customer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
		return
	}

	c.JSON(http.StatusOK, customer)
}

func DeleteCustomer(c *gin.Context) {
	var customer models.Customer
	if err := db.DB.First(&customer, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}

	if err := db.DB.Delete(&customer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete customer"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "customer deleted successfully"})
}
