package main

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/sergiosegrera/store/server/api"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := gin.Default()

	router.Use(gzip.Gzip(gzip.DefaultCompression))

	router.Static("/static", "./static/")

	router.GET("/api/products", api.GetProducts)
	router.GET("/api/product/:name", api.GetProduct)
	router.POST("/api/cart", api.PostCart)
	router.POST("/api/checkout", api.PostCheckout)
	// TODO:
	//	router.POST("/api/order/confirm", api.PostConfirm)
	// router.POST("/api/order/cancel")

	// Admin
	router.POST("/api/login", api.PostLogin)
	router.POST("/api/refresh", api.PostRefresh)

	// Dashboard
	router.GET("/api/admin/products", api.Auth(api.GetAllProducts))

	router.GET("/api/admin/product/:name", api.Auth(api.GetAllProduct))
	router.PUT("/api/admin/product", api.Auth(api.PutProduct))
	router.POST("/api/admin/product", api.Auth(api.PostProduct))

	router.POST("/api/admin/image", api.Auth(api.PostImage))

	router.GET("/api/admin/stocks/:name", api.Auth(api.GetStocks))
	router.PUT("/api/admin/stocks/:name", api.Auth(api.PutStock))
	router.POST("/api/admin/stocks/:name", api.Auth(api.PostStock))

	router.Run(":" + port)
}
