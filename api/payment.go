package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/plutov/paypal/v3"
	"github.com/sergiosegrera/store/db"
	"github.com/sergiosegrera/store/models"
)

func PostCheckout(c *gin.Context) {
	var cart models.Cart
	err := c.BindJSON(&cart)
	if err != nil {
		c.JSON(400, gin.H{"message": "Couldnt not parse order object"})
		return
	}
	// Verify Cart with database first
	checkedCart := db.CheckCart(cart)

	purchaseUnit := paypal.PurchaseUnitRequest{
		Amount: &paypal.PurchaseUnitAmount{
			Value:    fmt.Sprintf("%.2f", checkedCart.Total),
			Currency: "CAD",
		},
	}

	appContext := paypal.ApplicationContext{
		BrandName: "Store",
		ReturnURL: "https://localhost:8080/confirm",
		CancelURL: "https://localhost:8080/cancel"}

	response, err := pp.CreateOrder(paypal.OrderIntentCapture, []paypal.PurchaseUnitRequest{purchaseUnit}, nil, &appContext)
	if err != nil {
		c.JSON(500, gin.H{"message": "Internal Server Error"})
		return
	}
	for _, link := range response.Links {
		if link.Rel == "approve" {
			c.JSON(200, gin.H{"message": link.Href})
			return
		}
	}
	c.JSON(500, gin.H{"message": "Internal Server Error"})
}
