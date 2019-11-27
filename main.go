package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sergiosegrera/calida/api"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("./templates/*.tmpl")
	router.Static("/static", "./static")

	router.GET("/", api.Home)

	router.Run(":8080")
}
