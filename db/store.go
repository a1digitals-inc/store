package db

import "log"

type Product struct {
	Id          int      `json:"id"`
	Name        string   `json:"name"`
	Images      []string `json:"images"`
	Description string   `json:"description"`
	Price       int      `json:"price"`
	Discount    float32  `jsonn:"discount"`
}

type ProductThumbnail struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Thumbnail string `json:"thumbnail"`
	Soldout   bool   `json:"soldout"`
}

func GetProduct(id string) *Product {
	var product Product
	err := db.QueryRow("SELECT productid, name, description, price, ROUND((price * discount)::numeric, 2) FROM products WHERE productid=$1", id).Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.Discount)
	if err != nil {
		log.Println(err)
	}
	rows, err := db.Query("SELECT image FROM products INNER JOIN productimages USING(productid) WHERE productid=$1", id)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		var image string
		rows.Scan(&image)
		product.Images = append(product.Images, image)
	}

	return &product
}

func GetProducts() *[]ProductThumbnail {
	var products []ProductThumbnail

	rows, err := db.Query("SELECT productid, name, thumbnail FROM products")
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		var product ProductThumbnail
		rows.Scan(&product.Id, &product.Name, &product.Thumbnail)

		var quantity int
		err := db.QueryRow("SELECT SUM(quantity) FROM products INNER JOIN productstock USING(productid) GROUP BY productid HAVING productid=$1", product.Id).Scan(&quantity)
		if err != nil {
			log.Println(err)
		}

		if quantity > 0 {
			product.Soldout = false
		} else {
			product.Soldout = true
		}
		products = append(products, product)
	}
	return &products
}
