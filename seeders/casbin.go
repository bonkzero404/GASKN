package seeders

import (
	"errors"
	"go-starterkit-project/database/driver"
	"go-starterkit-project/domain/stores"

	"gorm.io/gorm"
)

func CreateCasbinPermission(db *gorm.DB) error {
	enforcer := driver.Enforcer

	var user stores.User
	var role stores.Role

	var email string = "bonkzero404@gmail.com"

	errUser := db.Take(&user, "email = ?", email).Error

	if errors.Is(errUser, gorm.ErrRecordNotFound) {
		return errUser
	}

	var roleName string = "Super Administrator"

	errRole := db.Take(&role, "role_name = ? AND role_type = ?", roleName, stores.SA).Error

	if errors.Is(errRole, gorm.ErrRecordNotFound) {
		return errRole
	}

	if g, err := enforcer.AddGroupingPolicy(user.ID.String(), role.ID.String(), "*"); !g {
		return err
	}

	if p, err := enforcer.AddPolicy(role.ID.String(), "*", "*", "POST|GET|UPDATE|DELETE"); !p {
		return err
	}

	return nil
}
