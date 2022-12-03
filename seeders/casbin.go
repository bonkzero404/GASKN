package seeders

import (
	"errors"
	"gaskn/config"
	"gaskn/database/stores"
	"gaskn/driver"
	"gorm.io/gorm"
)

func CreateCasbinPermission(db *gorm.DB) error {
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

	if g, err := driver.AddGroupingPolicy(
		user.ID.String(),
		role.ID.String(),
		"*",
		config.Config("ADMIN_FULLNAME"),
		config.Config("ADMIN_ROLENAME"),
		"",
	); !g {
		return err
	}

	if p, err := driver.AddPolicy(
		role.ID.String(),
		"*",
		"*",
		"GET|POST|PUT|DELETE",
		"",
		config.Config("ADMIN_ROLENAME"),
		"",
	); !p {
		return err
	}

	return nil
}
