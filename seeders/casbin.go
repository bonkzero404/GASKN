package seeders

import (
	"errors"
	"gaskn/config"
	"gaskn/database/driver"
	"gaskn/database/stores"

	"gorm.io/gorm"
)

func CreateCasbinPermission(db *gorm.DB) error {
	enforcer := driver.Enforcer

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

	if g, err := enforcer.AddGroupingPolicy(user.ID.String(), role.ID.String(), "*"); !g {
		return err
	}

	if p, err := enforcer.AddPolicy(role.ID.String(), "*", "*", "GET|POST|PUT|DELETE"); !p {
		return err
	}

	return nil
}
