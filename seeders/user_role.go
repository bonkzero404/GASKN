package seeders

import (
	"errors"
	"go-starterkit-project/config"
	"go-starterkit-project/domain/stores"

	"gorm.io/gorm"
)

func CreateUserRole(db *gorm.DB) error {
	var userRole stores.RoleUser
	var user stores.User
	var role stores.Role

	errUser := db.Take(&user, "email = ?", config.Config("ADMIN_EMAIL")).Error

	if errors.Is(errUser, gorm.ErrRecordNotFound) {
		return errUser
	}

	errRole := db.Take(&role, "role_name = ? AND role_type = ?", config.Config("ADMIN_ROLENAME"), stores.SA).Error

	if errors.Is(errRole, gorm.ErrRecordNotFound) {
		return errRole
	}

	err := db.Take(&userRole, "user_id = ? AND role_id = ?", user.ID.String(), role.ID.String()).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {

		userRole = stores.RoleUser{
			RoleId:   role.ID,
			UserId:   user.ID,
			IsActive: true,
		}

		return db.Create(&userRole).Error
	}

	return nil
}
