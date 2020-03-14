package db

type Product struct {
	Id          int      `json:"id"`
	Name        string   `json:"name"`
	Images      []string `json:"images"`
	Description string   `json:"description"`
	Price       int      `json:"price"`
	Discount    float32  `json:"discount"`
	Public      bool     `json:"public"`
}

type ProductThumbnail struct {
	Identifier string `json:"identifier"`
	Name       string `json:"name"`
	Thumbnail  string `json:"thumbnail"`
	Soldout    bool   `json:"soldout"`
}

func GetProduct(id string) *Product {
	var product Product
	db.QueryRow("SELECT productid, name, description, price, ROUND((price * discount)::numeric, 2), public FROM products WHERE identifier=$1 AND public=TRUE", id).Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.Discount, &product.Public)
	rows, _ := db.Query("SELECT image FROM products INNER JOIN productimages USING(productid) WHERE productid=$1", product.Id)
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

	rows, _ := db.Query("SELECT productid, name, thumbnail, identifier FROM products WHERE public=TRUE ORDER BY name")
	defer rows.Close()

	for rows.Next() {
		var product ProductThumbnail
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
