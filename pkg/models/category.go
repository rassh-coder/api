package models

type Category struct {
	Id   int    `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type CategoryInput struct {
	Name string `json:"name" binding:"required"`
}
