package services

import (
	"errors"
	"mvc-orm/internal/models"
	"mvc-orm/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(name, email, plainPassword string) (*models.User, error) {
	existingUser, _ := repository.FindUserByEmail(email)
	if existingUser != nil {
		return nil, errors.New("Email already exists")
	}

	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("Failed to hash password")
	}

	newUser := &models.User{
		Name: name,
		Email: email,
		Password:  string(hashedBytes),
	}
	if err := repository.CreateUser(newUser); err != nil {
		return nil, errors.New("Failed to save user to database")
	}

	return newUser, nil
}