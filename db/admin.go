package db

func InsertProduct(n string, p bool, t string, d string, pr string, di string, i string, im []string) error {
	var id int
	err := db.QueryRow("INSERT INTO products VALUES(DEFAULT, $1, $2, $3, $4, null, current_timestamp, $5, $6, $7) RETURNING productid", n, d, pr, di, t, p, i).Scan(&id)
	if err != nil {
		return err
	}
	for _, image := range im {
		_, err := db.Exec("INSERT INTO productimages VALUES($1, $2)", id, image)
		if err != nil {
			return err
		}
	}
	return err
}

func UpdateProduct(id int, n string, p bool, t string, d string, pr string, di string, i string, im []string) error {
	_, err := db.Exec("UPDATE products SET name=$2, description=$3, price=$4, discount=$5, thumbnail=$6, public=$7, identifier=$8 WHERE productid=$1", id, n, d, pr, di, t, p, i)
	if len(im) > 0 {
		// TODO: Maybe also delete images locally?
		_, err = db.Exec("DELETE FROM productimages WHERE productid=$1", id)
		for _, image := range im {
			_, err := db.Exec("INSERT INTO productimages VALUES($1, $2)", id, image)
			if err != nil {
				return err
			}
		}
	}
	return err
}
