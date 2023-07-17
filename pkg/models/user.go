package models

type User struct {
	Id           int    `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	UserName     string `json:"username"`
	PasswordHash string `json:"-"`
}

type UserInputFields struct {
	FirstName string `json:"first_name" binding:"required" db:"first_name"`
	LastName  string `json:"last_name" binding:"required" db:"last_name"`
	UserName  string `json:"username" binding:"required" db:"username"`
	Password  string `json:"password" binding:"required" db:"password_hash"`
}

type UserInputCreds struct {
	UserName     string `json:"username"`
	Password     string `json:"password"`
	CopyPassword string `json:"copy_password,omitempty"`
}
