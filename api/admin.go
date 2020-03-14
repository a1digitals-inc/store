package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sergiosegrera/store/db"
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
	name := c.PostForm("name")
	if name == "" {
		c.JSON(422, gin.H{"message": "Invalid name"})
		return
	}

	identifier := c.PostForm("identifier")
	// TODO: Verify that the identifier doesn't contain spaces or special characters
	if identifier == "" {
		c.JSON(422, gin.H{"message": "Invalid name"})
		return
	}

	publicCheck := c.PostForm("public")
	var public bool
	if publicCheck == "on" {
		public = true
	} else {
		public = false
	}

	// TODO: Limit thumbnail and images file size
	thumbnailFile, err := c.FormFile("thumbnail")
	if err != nil {
		c.JSON(422, gin.H{"message": "Invalid thumbnail"})
		return
	}
	thumbnailExtension := filepath.Ext(thumbnailFile.Filename)
	thumbnail := "/static/images/" + identifier + "-thumbnail" + thumbnailExtension
	err = c.SaveUploadedFile(thumbnailFile, "."+thumbnail)
	if err != nil {
		c.JSON(500, gin.H{"message": "Error saving thumbnail"})
		return
	}

	form, err := c.MultipartForm()
	imageFiles := form.File["images"]
	if err != nil {
		c.JSON(422, gin.H{"message": "Invalid images"})
		return
	}
	var images []string
	for i, imageFile := range imageFiles {
		imageExtension := filepath.Ext(imageFile.Filename)
		image := "/static/images/" + identifier + "-" + strconv.Itoa(i) + imageExtension
		err = c.SaveUploadedFile(imageFile, "."+image)
		if err != nil {
			c.JSON(500, gin.H{"message": "Error saving images"})
			return
		}
		images = append(images, image)
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

	// TODO: If identifier exists update instead of insert
	err = db.InsertProduct(name, public, thumbnail, description, price, discount, identifier, images)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Changes saved"})
}
