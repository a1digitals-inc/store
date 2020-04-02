package db

import "github.com/sergiosegrera/store/models"

func InsertProduct(product *models.Product) error {
	err := db.QueryRow(
		"INSERT INTO products (name, identifier, thumbnail, description, price, discount, created) VALUES ($1, $2, $3, $4, $5, $6, current_timestamp) RETURNING productid",
		product.Name,
		product.Identifier,
		product.Thumbnail,
		product.Description,
		product.Price,
		product.Discount,
	).Scan(&product.Id)
	if err != nil {
		return err
	}
	for _, image := range product.Images {
		_, err := db.Exec("INSERT INTO productimages VALUES($1, $2)", product.Id, image)
		if err != nil {
			return err
		}
	}
	return err
}

func UpdateProduct(identifier string, product *models.Product) error {
	err := db.QueryRow(
		"UPDATE products SET name=$2, description=$3, price=$4, discount=$5, thumbnail=$6, public=$7, identifier=$8 WHERE identifier=$1",
		identifier,
		product.Name,
		product.Description,
		product.Price,
		product.Discount,
		product.Thumbnail,
		product.Public,
		product.Identifier,
	).Scan(&product.Id)
	if len(product.Images) > 0 {
		_, err = db.Exec("DELETE FROM productimages WHERE productid=$1", product.Id)
		for _, image := range product.Images {
			_, err := db.Exec("INSERT INTO productimages VALUES($1, $2)", product.Id, image)
			if err != nil {
				return err
			}
		}
	}
	return err
}

func GetImages(identifier string) ([]string, error) {
	var images []string
	rows, err := db.Query("SELECT image FROM productimages INNER JOIN products USING (productid) WHERE identifier=$1", identifier)
	if err != nil {
		return images, err
	}
	defer rows.Close()

	for rows.Next() {
		var image string
		err = rows.Scan(&image)
		if err != nil {
			return images, err
		}

		images = append(images, image)
	}
	return images, err
}
