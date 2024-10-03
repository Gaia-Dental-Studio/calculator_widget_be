package model

import "gorm.io/gorm"

// Struct untuk menyimpan data produk
type Product struct {
	gorm.Model
	NameProduct  string `json:"name_product"`
	Description  string `json:"description"`
	Category     string `json:"category"`
	Image        string `json:"image"`
	Document     string `json:"document"`
	Price        string `json:"price"`
	FreeWarranty string `json:"free_warranty"`
}
