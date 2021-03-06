package db

import (
	"database/sql"
	"github.com/sergiosegrera/store/server/models"
)

func GetProduct(id string, p bool) (*models.Product, error) {
	var product models.Product
	var err error
	if p {
		err = db.QueryRow("SELECT productid, name, description, price, discount, public, thumbnail FROM products WHERE identifier=$1 AND public=TRUE", id).Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.Discount, &product.Public, &product.Thumbnail)
	} else {
		err = db.QueryRow("SELECT productid, name, description, price, discount, public, thumbnail FROM products WHERE identifier=$1", id).Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.Discount, &product.Public, &product.Thumbnail)
	}
	imageRows, _ := db.Query("SELECT image FROM productimages WHERE productid=$1", product.Id)
	defer imageRows.Close()

	for imageRows.Next() {
		var image string
		imageRows.Scan(&image)
		product.Images = append(product.Images, image)
	}

	optionRows, _ := db.Query("SELECT option FROM productstock WHERE productid=$1 AND quantity > 0", product.Id)
	defer optionRows.Close()

	for optionRows.Next() {
		var option string
		optionRows.Scan(&option)
		product.Options = append(product.Options, option)
	}

	return &product, err
}

func GetProductId(identifier string) (int, error) {
	var id int
	err := db.QueryRow("SELECT productid FROM products WHERE identifier=$1", identifier).Scan(&id)
	return id, err
}

func GetProducts(p bool) *[]models.ProductThumbnail {
	var products []models.ProductThumbnail

	var rows *sql.Rows
	if p {
		rows, _ = db.Query("SELECT productid, name, thumbnail, identifier FROM products WHERE public=TRUE ORDER BY name")
	} else {
		rows, _ = db.Query("SELECT productid, name, thumbnail, identifier FROM products ORDER BY created DESC")
	}
	defer rows.Close()

	for rows.Next() {
		var product models.ProductThumbnail
		var id int
		rows.Scan(&id, &product.Name, &product.Thumbnail, &product.Identifier)

		var quantity int
		db.QueryRow("SELECT SUM(quantity) FROM products INNER JOIN productstock USING(productid) GROUP BY productid HAVING productid=$1", id).Scan(&quantity)

		if quantity > 0 {
			product.Soldout = false
		} else {
			product.Soldout = true
		}
		products = append(products, product)
	}
	return &products
}
