package db

import "log"

func GetPassword() string {
	// TODO: Return error on no password found
	var password string
	err := db.QueryRow("SELECT value FROM settings WHERE name='password'").Scan(&password)
	if err != nil {
		log.Println(err)
	}

	return password
}

func SetPassword(p string) error {
	_, err := db.Exec("UPDATE settings SET value=$1 WHERE name='password'", p)
	return err
}
