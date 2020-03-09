package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"os"
)

var (
	db *sql.DB
)

func NewDatabase() (*sql.DB, error) {
	var err error

	url, _ := os.LookupEnv("DATABASE_URL")

	db, err = sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func Ping() error {
	err := db.Ping()
	return err
}
