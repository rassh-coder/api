package repository

import (
	"MyApi/pkg/models"
	"database/sql"
	"fmt"
)

type CategoriesRepository struct {
	db *sql.DB
}

func NewCategoriesRepository(db *sql.DB) *CategoriesRepository {
	return &CategoriesRepository{db: db}
}

func (r *CategoriesRepository) CreateCategory(input models.CategoryInput) (int64, error) {

	query := fmt.Sprintf("INSERT INTO %s(`name`) VALUES(?)", categoriesTable)
	row, err := r.db.Exec(query, input.Name)

	if err != nil {
		return 0, err
	}
	id, err := row.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *CategoriesRepository) GetAllCategories() ([]models.Category, error) {
	var data []models.Category
	rows, err := r.db.Query(fmt.Sprintf("SELECT * FROM %s", categoriesTable))
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var category models.Category
		err = rows.Scan(&category.Id, &category.Name)
		if err != nil {
			return nil, err
		}

		data = append(data, category)
	}

	return data, nil
}

func (r *CategoriesRepository) UpdateCategory(id int, input models.CategoryInput) (models.Category, error) {
	var category models.Category
	query := fmt.Sprintf("UPDATE %s SET name=? WHERE id=?", categoriesTable)
	_, err := r.db.Exec(query, input.Name, id)
	if err != nil {
		return category, err
	}

	err = r.db.QueryRow(fmt.Sprintf("SELECT * FROM %s WHERE id=?", categoriesTable), id).Scan(&category.Id, &category.Name)
	if err != nil {
		return category, err
	}

	return category, nil
}

func (r *CategoriesRepository) DeleteCategory(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=?", categoriesTable)
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *CategoriesRepository) GetCategory(id int) (models.Category, error) {
	var category models.Category

	err := r.db.QueryRow(fmt.Sprintf("SELECT * FROM %s WHERE id = ?", categoriesTable), id).Scan(&category.Id, &category.Name)
	if err != nil {
		return category, err
	}

	return category, nil
}
