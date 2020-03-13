package api

import (
	"github.com/gin-gonic/gin"
)

func Dashboard(c *gin.Context) {
	c.HTML(200, "index.tmpl", gin.H{
		"title":  "dashboard",
		"bundle": "dashboard",
	})
}

func AdminProducts(c *gin.Context) {
	c.HTML(200, "index.tmpl", gin.H{
		"title":  "store products",
		"bundle": "adminproducts",
	})
}
