package db

//import "log"

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
