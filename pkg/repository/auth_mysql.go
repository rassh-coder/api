package repository

import (
	"MyApi/pkg/models"
	"database/sql"
	"errors"
	"fmt"
)

type AuthRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) CreateUser(input models.UserInputFields) (int64, error) {
	query := fmt.Sprintf("INSERT INTO %s(first_name, last_name, username, password_hash) VALUES (?, ?, ?, ?)", usersTable)
	row, err := r.db.Exec(query, input.FirstName, input.LastName, input.UserName, input.Password)

	if err != nil {
		return 0, err
	}

	userId, err := row.LastInsertId()
	if err != nil {
		return 0, err
	}

	return userId, nil
}

func (r *AuthRepository) GetUser(creds models.UserInputCreds) (models.User, error) {
	var user models.User

	query := fmt.Sprintf("SELECT * FROM %s WHERE username=?", usersTable)
	err := r.db.QueryRow(query, creds.UserName).Scan(&user.Id, &user.UserName, &user.FirstName, &user.LastName, &user.PasswordHash)
	if err != nil {
		return models.User{}, err
	}
	if creds.Password != user.PasswordHash {
		return models.User{}, errors.New("wrong credentials")
	}

	return user, nil
}
