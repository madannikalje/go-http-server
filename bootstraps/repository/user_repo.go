package repository

import (
	"http-test/database"
	"http-test/models"
)

func FindUserByEmail(email string, user *models.User) error {
	if err := database.DB.Model(user).Where("email = ?", email).Select(); err != nil {
		return err
	}
	return nil
}

func CreateUser(user *models.User) error {
	if _, err := database.DB.Model(user).Insert(); err != nil {
		return err
	}
	return nil
}
