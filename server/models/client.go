package models

import (
	"regexp"
)

var (
	emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

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
	AddressLine1 string `json:"addressLine1"`
	AddressLine2 string `json:"addressLine2,omitempty"`
	City         string `json:"city"`
	State        string `json:"state"`
	Zip          string `json:"zip"`
}

func ValidateClient(client *Client) []string {
	var errs []string
	if client.Firstname == "" {
		errs = append(errs, "A firstname is required")
	}

	if client.Lastname == "" {
		errs = append(errs, "A lastname is required")
	}

	if client.Email == "" {
		errs = append(errs, "An email is required")
	} else if !emailRegex.MatchString(client.Email) {
		errs = append(errs, "Email is invalid")
	}

	if client.Phone == "" {
		errs = append(errs, "A phone number is required")
	}

	errs = append(errs, ValidateAddress(client.ShippingAddress)...)

	return errs
}

func ValidateAddress(address *Address) []string {
	var errs []string

	if address.Country == "" {
		errs = append(errs, "A country is required")
	} else if len(address.Country) != 2 {
		errs = append(errs, "Invalid Country code")
	}

	if address.AddressLine1 == "" {
		errs = append(errs, "An address line 1 is required")
	}

	if address.City == "" {
		errs = append(errs, "A city is required")
	}

	if address.State == "" {
		errs = append(errs, "A region is required")
	}

	if address.Zip == "" {
		errs = append(errs, "an area code is required")
	}

	return errs

}
