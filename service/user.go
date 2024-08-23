package service

import (
	"errors"
	"log"
	config "shopifun/Config"
	model "shopifun/Model"

	"gorm.io/gorm"
)

func InsertUser(registerUser *model.User) (*model.User, error) {
	db := config.DB
	// check existing user

	var existingUser model.User

	if err := db.Where("email = ?", registerUser.Email).First(&existingUser).Error; err == nil {
		return nil, errors.New("email already in use")
	}

	if err := db.Where("username = ?", registerUser.Username).First(&existingUser).Error; err == nil {
		return nil, errors.New("username already in use")
	}

	resultSave := db.Create(registerUser)

	if resultSave.Error != nil {
		log.Fatal(resultSave.Error)
		return registerUser, resultSave.Error
	}

	return registerUser, nil
}

func GetUsername(username string) (*model.User, error) {
	db := config.DB

	var existingUser model.User
	result := db.Where("username = ?", username).First(&existingUser)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	return &existingUser, nil
}
