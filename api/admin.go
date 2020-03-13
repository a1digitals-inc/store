package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sergiosegrera/store/db"
	"log"
	"path/filepath"
	"strconv"
)

func Dashboard(c *gin.Context) {
	c.HTML(200, "index.tmpl", gin.H{
		"title":  "dashboard",
		"bundle": "dashboard",
	})
}

func AdminProducts(c *gin.Context) {
	c.HTML(200, "index.tmpl", gin.H{
		"title":  "loading...",
		"bundle": "adminproducts",
	})
}

func AdminProduct(c *gin.Context) {
	c.HTML(200, "index.tmpl", gin.H{
		"title":  "loading...",
		"bundle": "adminproduct",
	})
}

func UpdateProduct(c *gin.Context) {
	id := c.Request.URL.Path
	log.Println(id)
	name := c.PostForm("name")
	if name == "" {
		c.JSON(422, gin.H{"error": "Invalid name"})
		return
	}

	identifier := c.PostForm("identifier")
	// TODO: Verify that the identifier doesn't contain spaces or special characters
	if identifier == "" {
		c.JSON(422, gin.H{"error": "Invalid name"})
		return
	}

	publicCheck := c.PostForm("public")
	var public bool
	if publicCheck == "on" {
		public = true
	} else {
		public = false
	}

	thumbnailFile, err := c.FormFile("thumbnail")
	if err != nil {
		c.JSON(422, gin.H{"error": "Invalid thumbnail"})
		return
	}
	thumbnailExtension := filepath.Ext(thumbnailFile.Filename)
	thumbnail := "/static/images/" + identifier + "-thumbnail" + thumbnailExtension
	err = c.SaveUploadedFile(thumbnailFile, thumbnail)
	if err != nil {
		log.Println(err)
		c.JSON(422, gin.H{"error": "Error saving file"})
	}

	// TODO: Save images
	form, err := c.MultipartForm()
	images := form.File["images"]
	if err != nil {
		c.JSON(422, gin.H{"error": "Invalid images"})
		return
	}
	for i, imageFile := range images {
		imageExtension := filepath.Ext(imageFile.Filename)
		image := "/static/images/" + identifier + "-" + strconv.Itoa(i) + imageExtension
		log.Println(image)
	}

	description := c.PostForm("description")

	price := c.PostForm("price")
	if price == "" {
		c.JSON(422, gin.H{"error": "Invalid price"})
		return
	}

	discount := c.PostForm("discount")
	if discount == "" {
		discount = "1"
	}

	err = db.InsertProduct(name, public, thumbnail, description, price, discount, identifier)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Changes saved"})
}
