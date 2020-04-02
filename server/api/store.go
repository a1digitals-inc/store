package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sergiosegrera/store/server/db"
)

func GetProducts(c *gin.Context) {
	products := db.GetProducts(true)
	c.JSON(200, gin.H{"message": products})
}

func GetProduct(c *gin.Context) {
	product, err := db.GetProduct(c.Param("name"), true)
	if err != nil {
		c.JSON(404, gin.H{"message": "Product not found"})
		return
	}
	c.JSON(200, gin.H{"message": product})
}
