package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sergiosegrera/store/server/db"
	"github.com/sergiosegrera/store/server/models"
)

func GetStocks(c *gin.Context) {
	identifier := c.Param("name")
	stocks, _ := db.GetStocks(identifier)
	c.JSON(200, gin.H{"message": stocks})
}

func PutStock(c *gin.Context) {
	var stock models.Stock
	err := c.BindJSON(&stock)
	if err != nil {
		c.JSON(422, gin.H{"message": "Could not parse stock"})
		return
	}

	identifier := c.Param("name")

	stock.ProductId, err = db.GetProductId(identifier)
	if err != nil {
		c.JSON(404, gin.H{"message": "Product not found"})
		return
	}

	err = db.UpdateStock(&stock)
	if err != nil {
		c.JSON(404, gin.H{"message": "No stock to update"})
	}

	c.JSON(200, gin.H{"message": "Stock updated"})
}

func PostStock(c *gin.Context) {
	var stock models.Stock
	err := c.BindJSON(&stock)
	if err != nil {
		c.JSON(422, gin.H{"message": "Could not parse stock"})
		return
	}

	identifier := c.Param("name")

	stock.ProductId, err = db.GetProductId(identifier)
	if err != nil {
		c.JSON(404, gin.H{"message": "Product not found"})
		return
	}

	err = db.InsertStock(&stock)
	if err != nil {
		c.JSON(500, gin.H{"message": "Internal server error"})
		return
	}

	c.JSON(200, gin.H{"message": "Stock created"})
}
