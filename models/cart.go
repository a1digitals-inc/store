package models

type Cart struct {
	Products      []CartProduct `json:"cartProducts"`
	PromotionCode string        `json:"promotionCode"`
	SubTotal      int64         `json:"subtotal"`
	Total         int64         `json:"total"`
}

type CartProduct struct {
	Name       string `json:"name"`
	Identifier string `json:"identifier"`
	Thumbnail  string `json:"thumbnail,omitempty"`
	Option     string `json:"option"`
	Quantity   int    `json:"quantity"`
	Price      int64  `json:"price,omitempty"`
}
