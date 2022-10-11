package seeders

import (
	"errors"
	"gaskn/database/stores"

	"gorm.io/gorm"
)

func CreateRole(db *gorm.DB, roleName string, roleDesc string, roleType string) error {
	var role stores.Role

	err := db.Take(&role, "role_name = ? AND role_type = ?", roleName, roleType).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {

		role = stores.Role{
			RoleName:        roleName,
			RoleDescription: roleDesc,
			RoleType:        roleType,
			IsActive:        true,
		}

		return db.Create(&role).Error
	}

	return nil
}
