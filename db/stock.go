package db

import (
	"errors"
	"github.com/sergiosegrera/store/models"
)

func GetStocks(identifier string) (*[]models.Stock, error) {
	var stocks []models.Stock

	rows, err := db.Query("SELECT option, quantity FROM products INNER JOIN productstock USING(productid) WHERE identifier=$1", identifier)

	defer rows.Close()
	for rows.Next() {
		var stock models.Stock
		rows.Scan(&stock.Option, &stock.Quantity)
		stocks = append(stocks, stock)
	}

	return &stocks, err
}

func InsertStock(id int, option string, quantity string) error {
	_, err := db.Exec("INSERT INTO productstock VALUES($1, $2, $3)", id, option, quantity)
	return err
}

func UpdateStock(id int, option string, quantity string) error {
	res, err := db.Exec("UPDATE productstock SET quantity=$3 WHERE productid=$1 AND option=$2", id, option, quantity)
	number, err := res.RowsAffected()
	if number == 0 {
		return errors.New("db: no row to update")
	}
	return err
}
