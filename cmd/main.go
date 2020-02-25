package main

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/sergiosegrera/calida/api"
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

	router.GET("/", api.Home)

	router.Run(":" + port)
}
