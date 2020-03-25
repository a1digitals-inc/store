package db

import (
	"github.com/sergiosegrera/store/models"
)

func StartOrder(order *models.Order) error {
	var clientId int
	err := db.QueryRow(
		"SELECT clientid FROM clients WHERE firstname=$1 AND lastname=$2 AND email=$3 AND phone=$4",
		order.OrderClient.Firstname,
		order.OrderClient.Lastname,
		order.OrderClient.Email,
		order.OrderClient.Phone,
	).Scan(&clientId)
	if err != nil {
		err = db.QueryRow(
			"INSERT INTO clients (firstname, lastname, email, phone, created) VALUES ($1, $2, $3, $4, current_timestamp) RETURNING clientid",
			order.OrderClient.Firstname,
			order.OrderClient.Lastname,
			order.OrderClient.Email,
			order.OrderClient.Phone,
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

	err = db.QueryRow("INSERT INTO orders (status, clientid, shippingid, created) VALUES('Created', $1, $2, current_timestamp) RETURNING orderid", clientId, addressId).Scan(&order.Id)
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
	err := db.QueryRow("SELECT orderid FROM orders WHERE orderid=$1 AND status='Created'", order.Id).Scan(&order.Id)
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
	var err error
	return err
}

func CancelOrder(order *models.Order) {

}
