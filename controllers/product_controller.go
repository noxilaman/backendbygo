package controllers

import (
	"example/web-service-gin/db"
	"example/web-service-gin/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
    var products []models.Product

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

    // Get total count of products
    var totalCount int64
    if err := db.DB.Model(&models.Product{}).Count(&totalCount).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products count"})
        return
    }

    // Fetch products with limit and offset
    if err := db.DB.Limit(pageSize).Offset(offset).Find(&products).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
        return
    }

    // Calculate total pages
    totalPages := (int(totalCount) + pageSize - 1) / pageSize

    // Prepare response
    response := gin.H{
        "data":       products,
        "page":       page,
        "pageSize":   pageSize,
        "total":      totalCount,
        "totalPages": totalPages,
    }

    c.JSON(http.StatusOK, response)
}

func GetProduct(c *gin.Context) {
	var product models.Product
	if err := db.DB.First(&product, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}

func CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Create(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	c.JSON(http.StatusCreated, product)
}

func UpdateProduct(c *gin.Context) {
	var product models.Product
	if err := db.DB.First(&product, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Save(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
		return
	}

	c.JSON(http.StatusOK, product)
}

func DeleteProduct(c *gin.Context) {
	var product models.Product
	if err := db.DB.First(&product, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	if err := db.DB.Delete(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
