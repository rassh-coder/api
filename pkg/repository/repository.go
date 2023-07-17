package repository

import (
	"MyApi/pkg/models"
	"database/sql"
)

type Authorization interface {
	CreateUser(input models.UserInputFields) (int64, error)
	GetUser(creds models.UserInputCreds) (models.User, error)
}

type Categories interface {
	CreateCategory(input models.CategoryInput) (int64, error)
	GetAllCategories() ([]models.Category, error)
	UpdateCategory(id int, input models.CategoryInput) (models.Category, error)
	DeleteCategory(id int) error
	GetCategory(id int) (models.Category, error)
}

type Products interface {
	CreateProduct(input models.ProductInputFields) (int64, error)
}

type Repository struct {
	Authorization
	Categories
	Products
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: NewAuthRepository(db),
		Categories:    NewCategoriesRepository(db),
		Products:      NewProductsRepository(db),
	}
}
