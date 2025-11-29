package models

type User struct {
	ID       int    `json:"id" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Age      int    `json:"age" validate:"min=10"`
}
