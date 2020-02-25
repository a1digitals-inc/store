package api

import (
	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	c.HTML(200, "index.tmpl", gin.H{
		"title":  "home",
		"bundle": "home",
	})
}
