package db

//import "log"

func InsertProduct(n string, p bool, t string, d string, pr string, di string, i string) error {
	_, err := db.Exec("INSERT INTO products VALUES(DEFAULT, $1, $2, $3, $4, null, current_timestamp, $5, $6, $7)", n, d, pr, di, t, p, i)
	if err != nil {
		return err
	}
	//	id, err := response.LastInsertId()
	//	log.Println(id)
	return err
}
