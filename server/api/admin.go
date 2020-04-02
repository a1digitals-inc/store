package api

import (
	"github.com/gin-gonic/gin"
	"github.com/lithammer/shortuuid"
	"github.com/sergiosegrera/store/server/db"
	"github.com/sergiosegrera/store/server/models"
	"log"
	"os"
)

func GetAllProducts(c *gin.Context) {
	// Gets all products including private products
	products := db.GetProducts(false)
	c.JSON(200, gin.H{"message": products})
}

func GetAllProduct(c *gin.Context) {
	// Gets a product even if its private
	product, err := db.GetProduct(c.Param("name"), false)
	if err != nil {
		c.JSON(404, gin.H{"message": "Product not found"})
		return
	}
	c.JSON(200, gin.H{"message": product})
}

func PutProduct(c *gin.Context) {
	var product models.Product
	err := c.BindJSON(&product)
	if err != nil {
		c.JSON(422, gin.H{"message": "Could not parse product"})
	}

	name := c.Param("name")

	errs := models.ValidateProduct(&product)
	if len(errs) > 0 {
		c.JSON(422, gin.H{"message": errs})
		return
	}

	oldImages, _ := db.GetImages(name)
	for _, image := range oldImages {
		os.Remove(image)
	}

	err = db.UpdateProduct(name, &product)
	if err != nil {
		c.JSON(422, gin.H{"message": "Could not update"})
		return
	}

	c.JSON(200, gin.H{"message": "Product updated"})
}

func PostProduct(c *gin.Context) {
	var product models.Product
	err := c.BindJSON(&product)
	if err != nil {
		c.JSON(422, gin.H{"message": "Could not parse product"})
		return
	}

	errs := models.ValidateProduct(&product)
	if len(errs) > 0 {
		c.JSON(422, gin.H{"message": errs})
		return
	}

	err = db.InsertProduct(&product)
	if err != nil {
		c.JSON(500, gin.H{"message": "Internal server error"})
		return
	}

	c.JSON(200, gin.H{"message": "Product created"})
}

func PostImage(c *gin.Context) {
	id := shortuuid.New()
	dest := "/static/images/" + id + ".jpg"
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(422, gin.H{"message": "Invalid name"})
		return
	}

	err = CompressAndSaveFile(file, "."+dest)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{"message": "Internal server error"})
		return
	}

	c.JSON(200, gin.H{"message": dest})
}
