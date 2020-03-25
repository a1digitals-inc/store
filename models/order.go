package models

type Order struct {
	Id              int     `json:"id,omitempty"`
	PaymentId       string  `json:"paymentId"`
	PaymentMethodId string  `json:"paymentMethodId"`
	OrderClient     *Client `json:"client"`
	OrderCart       *Cart   `json:"cart"`
	Status          string  `json:"status,omitempty"`
}
