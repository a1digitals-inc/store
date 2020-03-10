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

	router.LoadHTMLGlob("./views/*.tmpl")
	router.Static("/static", "./static")

	// API
	router.GET("/api/products", api.GetProducts)
	router.GET("/api/product/:name", api.GetProduct)

	// HTML
	router.GET("/", api.Home)
	router.GET("/webstore", api.Store)
	router.GET("/product/:name", api.Product)

	router.Run(":" + port)
}
