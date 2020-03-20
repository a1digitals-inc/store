package models

type Product struct {
	Id          int      `json:"id"`
	Name        string   `json:"name"`
	Thumbnail   string   `json:"thumbnail"`
	Images      []string `json:"images"`
	Description string   `json:"description"`
	Options     []string `json:"options"`
	Price       int      `json:"price"`
	Discount    float32  `json:"discount"`
	Public      bool     `json:"public"`
}
