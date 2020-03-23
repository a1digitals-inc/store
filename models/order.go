package models

type Order struct {
	Id          int     `json:"id,omitempty"`
	Token       string  `json:"token"`
	OrderClient *Client `json:"client"`
	OrderCart   *Cart   `json:"cart"`
	Status      string  `json:"status,omitempty"`
}
