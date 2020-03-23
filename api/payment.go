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
		c.JSON(400, gin.H{"message": "Could not parse cart object"})
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
		CancelURL: "https://localhost:8080/cancel",
	}

	response, err := pp.CreateOrder(paypal.OrderIntentCapture, []paypal.PurchaseUnitRequest{purchaseUnit}, nil, &appContext)
	if err != nil {
		c.JSON(500, gin.H{"message": "Internal Server Error"})
		return
	}
	for _, link := range response.Links {
		if link.Rel == "approve" {
			err = db.StartOrder(&models.Order{Token: response.ID, OrderCart: checkedCart})
			if err != nil {
				c.JSON(500, gin.H{"message": "Internal Server Error"})
				return
			}
			c.JSON(200, gin.H{"message": link.Href})
			return
		}
	}
	c.JSON(500, gin.H{"message": "Internal server error"})
}

func PostConfirm(c *gin.Context) {
	var order models.Order
	err := c.BindJSON(&order)
	if err != nil {
		c.JSON(400, gin.H{"message": "Could not parse token"})
		return
	}

	paypalOrder, err := pp.GetOrder(order.Token)
	if err != nil {
		c.JSON(400, gin.H{"message": "Wrong token"})
		return
	}

	if paypalOrder.Status != "APPROVED" {
		c.JSON(400, gin.H{"message": "Payment not approved"})
		return
	}

	// TODO: Save client information
	err = db.ReserveOrder(&order)
	if err != nil {
		c.JSON(500, gin.H{"message": "Internal server error"})
		return
	}

	capture, err := pp.CaptureOrder(order.Token, paypal.CaptureOrderRequest{})
	if err != nil {
		c.JSON(500, gin.H{"message": "Failed capture"})
		return
	}

	order.OrderClient = &models.Client{
		Firstname: capture.Payer.Name.Surname,
		Lastname:  capture.Payer.Name.GivenName,
		Email:     capture.Payer.EmailAddress,
		//		Phone:     capture.Payer.Phone.PhoneNumber.NationalNumber,
		ShippingAddress: &models.Address{
			Country:      capture.PurchaseUnits[0].Shipping.Address.CountryCode,
			AddressLine1: capture.PurchaseUnits[0].Shipping.Address.AddressLine1,
			AddressLine2: capture.PurchaseUnits[0].Shipping.Address.AddressLine2,
			City:         capture.PurchaseUnits[0].Shipping.Address.AdminArea2,
			State:        capture.PurchaseUnits[0].Shipping.Address.AdminArea1,
			Zip:          capture.PurchaseUnits[0].Shipping.Address.PostalCode,
		},
	}
	err = db.CompleteOrder(&order)
	if err != nil {
		c.JSON(500, gin.H{"message": "Failed to save client information"})
		// TODO: Re-add reserved stock
		return
	}

	c.JSON(200, gin.H{"message": "Completed"})
}
