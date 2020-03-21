package models

type Cart struct {
	Products      []CartProduct `json:"cartProducts"`
	PromotionCode string        `json:"promotionCode"`
	Total         float32       `json:"total"`
}

type CartProduct struct {
	Identifier string `json:"identifier"`
	Option     string `json:"option"`
	Quantity   int    `json:"quantity"`
}
