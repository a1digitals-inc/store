package models

type ProductThumbnail struct {
	Identifier string `json:"identifier"`
	Name       string `json:"name"`
	Thumbnail  string `json:"thumbnail"`
	Soldout    bool   `json:"soldout"`
}
