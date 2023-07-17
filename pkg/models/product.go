package models

type Product struct {
	Id          int64   `json:"id"`
	SKU         string  `json:"sku"`
	Name        string  `json:"name"`
	Body        string  `json:"body"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Color       string  `json:"color"`
	Size        int     `json:"size"`
	Count       int     `json:"count"`
}

type ProductInputFields struct {
	SKU         string `json:"sku"`
	Name        string `json:"name" binding:"required"`
	Body        string `json:"body"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Color       string `json:"color"`
	Size        int    `json:"size"`
	Count       int    `json:"count"`
}
