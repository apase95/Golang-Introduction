package repository

import (
	"mvc-orm/internal/config"
	"mvc-orm/internal/models"
)

func CreateUser(user *models.User) error {
	result := config.DB.Create(user)
	return result.Error
}

func FindUserByEmail(email string) (*models.User, error) {
	var user models.User
	
	result := config.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	
	return &user, nil
}