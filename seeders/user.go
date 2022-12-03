package seeders

import (
	"errors"
	"gaskn/config"
	"gaskn/database/stores"
	"gaskn/utils"

	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB) error {
	var user stores.User

	err := db.Take(&user, "email = ?", config.Config("ADMIN_EMAIL")).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		hashPassword, _ := utils.HashPassword(config.Config("ADMIN_PASSWORD"))

		user = stores.User{
			FullName: config.Config("ADMIN_FULLNAME"),
			Email:    config.Config("ADMIN_EMAIL"),
			Phone:    config.Config("ADMIN_PHONE"),
			Password: hashPassword,
			IsActive: true,
		}

		return db.Create(&user).Error
	}

	return nil
}
