package services

import (
	"CRUD/internal/models"
	"errors"
)

var users = []models.User{
	{ID: 1, Name: "Olamide", Email: "olamide@example.com"},
	{ID: 2, Name: "John", Email: "john@example.com"},
}

var lastID = 2

func GetAllUsers() []models.User {
	return users
}

func GetUserByID(id int) (models.User, error) {
	for _, user := range users {
		if user.ID == id {
			return user, nil
		}
	}
	return models.User{}, errors.New("user not found")
}

func CreateUser(user models.User) models.User {
	lastID++
	user.ID = lastID
	users = append(users, user)
	return user
}

func UpdateUser(id int, updated models.User) (models.User, error) {
	for i, user := range users {
		if user.ID == id {
			users[i].Name = updated.Name
			users[i].Email = updated.Email
			return users[i], nil
		}
	}
	return models.User{}, errors.New("user not found")
}

func DeleteUser(id int) error {
	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return errors.New("user not found")
}
