package models

type Client struct {
	Id              string   `json:"id,omitempty"`
	Firstname       string   `json:"firstname"`
	Lastname        string   `json:"lastname"`
	Email           string   `json:"email"`
	Phone           string   `json:"phone"`
	ShippingAddress *Address `json:"shippingAddress,omitempty"`
	BillingAddress  *Address `json:"billingAddress,omitempty"`
}

type Address struct {
	Country      string `json:"country"`
	AddressLine1 string `json:"adressLine1"`
	AddressLine2 string `json:"adressLine2"`
	City         string `json:"city"`
	State        string `json:"state"`
	Zip          string `json:"zip"`
}
