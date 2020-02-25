package main

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/sergiosegrera/calida/api"
)

func main() {
	router := gin.Default()

	router.Use(gzip.Gzip(gzip.DefaultCompression))

	router.LoadHTMLGlob("./views/*.tmpl")
	router.Static("/static", "./static")

	router.GET("/", api.Home)

	router.Run(":8080")
}
