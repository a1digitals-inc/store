package db

import (
	"errors"
	"github.com/sergiosegrera/store/models"
)

func StartOrder(order *models.Order) error {
	err := db.QueryRow("INSERT INTO orders (token, status, created) VALUES($1, 'Created', current_timestamp) RETURNING orderid", order.Token).Scan(&order.Id)
	if err != nil {
		return err
	}
	for _, product := range order.OrderCart.Products {
		var productId int
		err = db.QueryRow("SELECT productid FROM products WHERE identifier=$1", product.Identifier).Scan(&productId)
		if err != nil {
			return err
		}
		_, err := db.Exec(
			"INSERT INTO orderitems (orderid, productid, option, quantity) VALUES($1, $2, $3, $4)",
			order.Id,
			productId,
			product.Option,
			product.Quantity,
		)
		if err != nil {
			return err
		}
	}
	return err
}

func ReserveOrder(order *models.Order) error {
	err := db.QueryRow("SELECT orderid FROM orders WHERE token=$1 AND status='Created'", order.Token).Scan(&order.Id)
	if err != nil {
		return err
	}

	rows, err := db.Query("SELECT productid, quantity, option FROM orderitems WHERE orderid=$1", order.Id)
	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		var productid int
		var quantity int
		var option string
		rows.Scan(&productid, &quantity, &option)
		_, err := db.Exec(
			"UPDATE productstock SET quantity=quantity-$2 WHERE quantity-$2 >= 0 AND option=$3 AND quantity != 0 AND productid=$1",
			productid,
			quantity,
			option,
		)
		if err != nil {
			return err
		}
	}

	_, err = db.Exec("UPDATE orders SET status='Reserved' WHERE orderid=$1", order.Id)

	return err
}

func CompleteOrder(order *models.Order) error {
	var clientId int
	err := db.QueryRow(
		"SELECT clientid FROM clients WHERE firstname=$1 AND lastname=$2 AND email=$3",
		order.OrderClient.Firstname,
		order.OrderClient.Lastname,
		order.OrderClient.Email,
	).Scan(&clientId)
	if err != nil {
		err = db.QueryRow(
			"INSERT INTO clients (firstname, lastname, email, created) VALUES ($1, $2, $3, current_timestamp) RETURNING clientid",
			order.OrderClient.Firstname,
			order.OrderClient.Lastname,
			order.OrderClient.Email,
		).Scan(&clientId)
		if err != nil {
			return err
		}
	}

	var addressId int
	err = db.QueryRow(
		"INSERT INTO addresses (clientid, country, addressline1, addressline2, city, zip, state) VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING addressid",
		clientId,
		order.OrderClient.ShippingAddress.Country,
		order.OrderClient.ShippingAddress.AddressLine1,
		order.OrderClient.ShippingAddress.AddressLine2,
		order.OrderClient.ShippingAddress.City,
		order.OrderClient.ShippingAddress.Zip,
		order.OrderClient.ShippingAddress.State,
	).Scan(&addressId)

	if err != nil {
		return err
	}

	result, err := db.Exec(
		"UPDATE orders SET status='Completed', clientid=$3, shippingid=$4 WHERE orderid=$1 AND token=$2 AND status='Reserved'",
		order.Id,
		order.Token,
		clientId,
		addressId,
	)
	count, err := result.RowsAffected()
	if count != 1 || err != nil {
		return errors.New("db: Could not update order with provided id and token")
	}

	return err
}

func CancelOrder(order *models.Order) {

}
