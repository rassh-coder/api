package repository

import (
	"MyApi/pkg/models"
	"database/sql"
	"fmt"
)

type ProductsRepository struct {
	db *sql.DB
}

func NewProductsRepository(db *sql.DB) *ProductsRepository {
	return &ProductsRepository{db: db}
}

func (r *ProductsRepository) CreateProduct(input models.ProductInputFields) (int64, error) {
	query := fmt.Sprintf("INSERT INTO %s (sku, name, body, description, price, color, size, count) VALUES(?,?,?,?,?,?,?,?)", productsTable)
	row, err := r.db.Exec(query, input.SKU, input.Name, input.Body, input.Description, input.Price, input.Color, input.Size, input.Count)
	if err != nil {
		return 0, err
	}

	id, err := row.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}
