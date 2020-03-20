package models

type Stock struct {
	ProductId int    `json:"productid"`
	Option    string `json:"option"`
	Quantity  int    `json:"quantity"`
}
