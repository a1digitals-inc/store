package db

import (
	"github.com/sergiosegrera/store/models"
)

func CheckCart(cart *models.Cart) *models.Cart {
	var checkedCart models.Cart
	for _, product := range cart.Products {
		if product.Quantity != 0 {
			var total int64
			err := db.QueryRow("SELECT name, thumbnail, price * discount * 100, (price * discount * $3 * 100)::integer FROM products INNER JOIN productstock USING(productid) WHERE identifier=$1 AND option=$2 AND quantity >= $3 AND quantity != 0 AND public='true'", product.Identifier, product.Option, product.Quantity).Scan(&product.Name, &product.Thumbnail, &product.Price, &total)
			if err == nil {
				checkedCart.Products = append(checkedCart.Products, product)
				checkedCart.SubTotal += total
			}
		}
	}
	checkedCart.Total = checkedCart.SubTotal
	if cart.PromotionCode != "" {
		var modifier float32
		err := db.QueryRow("SELECT modifier FROM promotions WHERE code=$1", cart.PromotionCode).Scan(&modifier)
		if err == nil {
			checkedCart.Total = int64(float32(checkedCart.Total) * modifier)
			checkedCart.PromotionCode = cart.PromotionCode
		}
	}
	return &checkedCart
}
