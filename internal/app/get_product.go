package app

import (
	"github.com/efmatui/product-api/internal/service"
	"github.com/gin-gonic/gin"
)

func GetAllProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		products, err := service.GetAllProduct()
		if err != nil {
			c.JSON(500, gin.H{"status": "Internal Server Error", "error": err})
		} else {
			c.JSON(200, gin.H{"status": "OK", "data": products})
		}

	}
}
