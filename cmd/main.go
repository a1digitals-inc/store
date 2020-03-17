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
	router.PUT("/api/product", api.Auth(api.UpdateProduct))
	router.POST("/api/login", api.PostLogin)
	router.GET("/api/admin/products", api.Auth(api.GetAllProducts))
	router.GET("/api/admin/product/:name", api.Auth(api.GetAllProduct))

	// Pages
	router.GET("/", api.Home)
	router.GET("/webstore", api.Store)
	router.GET("/product/:name", api.Product)

	// Admin Pages
	router.GET("/login", api.Login)

	router.GET("/dashboard", api.Auth(api.Dashboard))
	router.GET("/dashboard/products", api.Auth(api.AdminProducts))
	router.GET("/dashboard/product", api.Auth(api.AdminProduct))
	router.GET("/dashboard/product/:name", api.Auth(api.AdminProduct))

	router.Run(":" + port)
}
