package main

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/sergiosegrera/store/api"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := gin.Default()

	router.Use(gzip.Gzip(gzip.DefaultCompression))

	router.Static("/static", "./static")

	router.GET("/api/products", api.GetProducts)
	router.GET("/api/product/:name", api.GetProduct)
	router.POST("/api/cart", api.PostCart)
	// TODO
	router.POST("/api/checkout", api.PostCheckout)
	router.POST("/api/order/confirm", api.PostConfirm)
	// router.POST("/api/order/cancel")

	// Admin
	router.POST("/api/login", api.PostLogin)
	router.PUT("/api/product", api.Auth(api.UpdateProduct))

	// Dashboard
	router.GET("/api/admin/products", api.Auth(api.GetAllProducts))
	router.GET("/api/admin/product/:name", api.Auth(api.GetAllProduct))
	router.GET("/api/admin/stocks/:name", api.Auth(api.GetStocks))
	router.PUT("/api/admin/stock/:name", api.Auth(api.PutStock))

	router.Run(":" + port)
}
