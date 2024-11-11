package module

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Id          int      `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Picture     string   `json:"picture"`
	Price       float32  `json:"price"`
	Categories  []string `json:"categories"`
}
