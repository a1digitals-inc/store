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

func GetProducts(c *gin.Context) {
	products := db.GetProducts()
	c.JSON(200, products)
}

func GetProduct(c *gin.Context) {
	product := db.GetProduct(c.Param("name"))
	c.JSON(200, product)
}
