package models

import "strings"

type Product struct {
	Id          int      `json:"id"`
	Name        string   `json:"name"`
	Identifier  string   `json:"identifier,omitempty"`
	Thumbnail   string   `json:"thumbnail"`
	Images      []string `json:"images"`
	Description string   `json:"description"`
	Options     []string `json:"options"`
	Price       int      `json:"price"`
	Discount    float32  `json:"discount"`
	Public      bool     `json:"public"`
}

func ValidateProduct(product *Product) []string {
	var errs []string

	if product.Name == "" {
		errs = append(errs, "A name is required")
	}

	if product.Identifier == "" || strings.ContainsAny(product.Identifier, " /?$#") {
		errs = append(errs, "A valid identifier is required")
	}

	if product.Thumbnail == "" {
		errs = append(errs, "A thumbnail is required")
	}

	if len(product.Images) == 0 {
		errs = append(errs, "Atleast one image is required")
	}

	return errs
}
