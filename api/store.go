package api

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/sergiosegrera/store/db"
)

var (
	dbc *sql.DB
)

func init() {
	var err error
	dbc, err = db.NewDatabase()

	if err != nil {
		panic(err)
	}
}

func Store(c *gin.Context) {
	c.HTML(200, "index.tmpl", gin.H{
		"title":  "store",
		"bundle": "store",
	})
}

func Product(c *gin.Context) {
	c.HTML(200, "index.tmpl", gin.H{
		"title":  "loading...",
		"bundle": "product",
	})
}

func GetProducts(c *gin.Context) {
	products := db.GetProducts(true)
	c.JSON(200, products)
}

func GetProduct(c *gin.Context) {
	product, err := db.GetProduct(c.Param("name"), true)
	if err != nil {
		c.JSON(404, gin.H{"message": "Product not found"})
		return
	}
	c.JSON(200, product)
}
