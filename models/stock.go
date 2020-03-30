package models

type Stock struct {
	ProductId int    `json:"productid,omitempty"`
	Option    string `json:"option"`
	Quantity  int    `json:"quantity"`
}
