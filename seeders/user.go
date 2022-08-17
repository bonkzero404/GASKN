package seeders

import (
	"errors"
	"go-starterkit-project/domain/stores"
	"go-starterkit-project/utils"

	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB) error {
	var user stores.User
	var email string = "bonkzero404@gmail.com"
	var password string = "janitrapanji"

	err := db.Take(&user, "email = ?", email).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		hashPassword, _ := utils.HashPassword(password)

		user = stores.User{
			FullName: "Janitra Panji",
			Email:    email,
			Phone:    "+6281299579761",
			Password: hashPassword,
			IsActive: true,
		}

		return db.Create(&user).Error
	}

	return nil
}
