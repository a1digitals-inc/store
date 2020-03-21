package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sergiosegrera/store/db"
	"github.com/sergiosegrera/store/models"
)

func PostCart(c *gin.Context) {
	var cart models.Cart
	err := c.BindJSON(&cart)
	if err != nil {
		c.JSON(400, gin.H{"message": "Could not parse cart object"})
		return
	}
	checkedCart := db.CheckCart(cart)
	c.JSON(200, gin.H{"message": checkedCart})
}
