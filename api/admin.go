package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sergiosegrera/store/db"
	"path/filepath"
	"strconv"
	"strings"
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

func AdminProduct(c *gin.Context) {
	c.HTML(200, "index.tmpl", gin.H{
		"title":  "new product",
		"bundle": "adminproduct",
	})
}

func GetAllProducts(c *gin.Context) {
	// Gets all products including private products
	products := db.GetProducts(false)
	c.JSON(200, products)
}

func GetAllProduct(c *gin.Context) {
	// Gets a product even if its private
	product, err := db.GetProduct(c.Param("name"), false)
	if err != nil {
		c.JSON(404, gin.H{"message": "Product not found"})
		return
	}
	c.JSON(200, product)
}

func UpdateProduct(c *gin.Context) {
	name := c.PostForm("name")
	if name == "" {
		c.JSON(422, gin.H{"message": "Invalid name"})
		return
	}

	identifier := c.PostForm("identifier")
	if identifier == "" || strings.ContainsAny(identifier, " /?$#") {
		c.JSON(422, gin.H{"message": "Invalid identifier"})
		return
	}

	product, err := db.GetProduct(identifier, false)
	exists := true
	if err != nil {
		exists = false
	}

	publicCheck := c.PostForm("public")
	var public bool
	if publicCheck == "on" {
		public = true
	} else {
		public = false
	}

	// TODO: Limit thumbnail and images file size
	thumbnailFile, err := c.FormFile("thumbnailFile")
	var thumbnailPath string
	if err != nil {
		if !exists {
			c.JSON(422, gin.H{"message": "Invalid thumbnail"})
			return
		}
	} else {
		thumbnailExtension := filepath.Ext(thumbnailFile.Filename)
		thumbnailPath = "/static/images/" + identifier + "-thumbnail" + thumbnailExtension
		err = c.SaveUploadedFile(thumbnailFile, "."+thumbnailPath)
		if err != nil {
			c.JSON(500, gin.H{"message": "Error saving thumbnail"})
			return
		}
		product.Thumbnail = thumbnailPath
	}

	form, err := c.MultipartForm()
	imageFiles := form.File["imageFiles"]
	var imagePaths []string
	if err != nil {
		if !exists {
			c.JSON(422, gin.H{"message": "Invalid images"})
			return
		}
	} else {
		for i, imageFile := range imageFiles {
			imageExtension := filepath.Ext(imageFile.Filename)
			image := "/static/images/" + identifier + "-" + strconv.Itoa(i) + imageExtension
			err = c.SaveUploadedFile(imageFile, "."+image)
			if err != nil {
				c.JSON(500, gin.H{"message": "Error saving images"})
				return
			}
			imagePaths = append(imagePaths, image)
		}
		product.Images = imagePaths
	}

	description := c.PostForm("description")

	price := c.PostForm("price")
	if price == "" {
		c.JSON(422, gin.H{"message": "Invalid price"})
		return
	}

	discount := c.PostForm("discount")
	if discount == "" {
		discount = "1"
	}

	if !exists {
		err = db.InsertProduct(name, public, thumbnailPath, description, price, discount, identifier, imagePaths)
		if err != nil {
			c.JSON(500, gin.H{"message": err.Error()})
			return
		}
	} else {
		err = db.UpdateProduct(product.Id, name, public, product.Thumbnail, description, price, discount, identifier, product.Images)
		if err != nil {
			c.JSON(500, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(200, gin.H{"message": "Changes saved"})
}
