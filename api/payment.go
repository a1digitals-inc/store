package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sergiosegrera/store/db"
	"github.com/sergiosegrera/store/models"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/paymentintent"
	"log"
	"strconv"
)

func PostCheckout(c *gin.Context) {
	var order models.Order
	err := c.BindJSON(&order)
	if err != nil {
		c.JSON(400, gin.H{"message": "Could not parse cart object"})
		return
	}

	errs := models.ValidateClient(order.OrderClient)
	if len(errs) > 0 {
		c.JSON(422, gin.H{"message": errs})
		return
	}

	order.OrderCart = db.CheckCart(order.OrderCart)

	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(order.OrderCart.Total),
		Currency: stripe.String(string(stripe.CurrencyCAD)),
		PaymentMethodTypes: stripe.StringSlice([]string{
			"card",
		}),
		ReceiptEmail: stripe.String(order.OrderClient.Email),
		Shipping: &stripe.ShippingDetailsParams{
			Address: &stripe.AddressParams{
				City:       stripe.String(order.OrderClient.ShippingAddress.City),
				Country:    stripe.String(order.OrderClient.ShippingAddress.Country),
				Line1:      stripe.String(order.OrderClient.ShippingAddress.AddressLine1),
				Line2:      stripe.String(order.OrderClient.ShippingAddress.AddressLine2),
				PostalCode: stripe.String(order.OrderClient.ShippingAddress.Zip),
				State:      stripe.String(order.OrderClient.ShippingAddress.State),
			},
			Name:  stripe.String(order.OrderClient.Firstname + " " + order.OrderClient.Lastname),
			Phone: stripe.String(order.OrderClient.Phone),
		},
	}

	err = db.StartOrder(&order)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{"message": "Internal Server Error"})
		return
	}

	params.AddMetadata("orderid", strconv.Itoa(order.Id))

	intent, err := paymentintent.New(params)
	if err != nil {
		c.JSON(500, gin.H{"message": "Internal Server Error"})
		return
	}

	c.JSON(200, gin.H{"message": intent.ClientSecret})
}
