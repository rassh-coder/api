package service

import (
	"MyApi/pkg/models"
	"MyApi/pkg/repository"
)

type ProductsService struct {
	repos repository.Products
}

func NewProductsService(repos repository.Products) *ProductsService {
	return &ProductsService{repos: repos}
}

func (s ProductsService) CreateProduct(input models.ProductInputFields) (int64, error) {
	return s.repos.CreateProduct(input)
}
