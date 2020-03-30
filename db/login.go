package db

import "log"

func GetPassword() string {
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
