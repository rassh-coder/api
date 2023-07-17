package service

import (
	"MyApi/pkg/models"
	"MyApi/pkg/repository"
)

type Authorization interface {
	CreateUser(input models.UserInputFields) (int64, error)
	SignIn(creds models.UserInputCreds) (string, error)
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

type Service struct {
	Authorization
	Categories
	Products
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Categories:    NewCategoriesService(repos.Categories),
		Products:      NewProductsService(repos.Products),
	}
}
