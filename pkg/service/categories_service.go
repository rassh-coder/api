package service

import (
	"MyApi/pkg/models"
	"MyApi/pkg/repository"
)

type CategoriesService struct {
	repos repository.Categories
}

func NewCategoriesService(repos repository.Categories) *CategoriesService {
	return &CategoriesService{repos: repos}
}

func (s *CategoriesService) CreateCategory(input models.CategoryInput) (int64, error) {
	return s.repos.CreateCategory(input)
}

func (s *CategoriesService) GetAllCategories() ([]models.Category, error) {
	return s.repos.GetAllCategories()
}

func (s *CategoriesService) UpdateCategory(id int, input models.CategoryInput) (models.Category, error) {
	return s.repos.UpdateCategory(id, input)
}

func (s *CategoriesService) DeleteCategory(id int) error {
	return s.repos.DeleteCategory(id)
}

func (s *CategoriesService) GetCategory(id int) (models.Category, error) {
	return s.repos.GetCategory(id)
}
