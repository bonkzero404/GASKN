package database

import (
	"gaskn/database/driver"
	"gaskn/database/stores"
)

// MigrateDB /*
func MigrateDB() {
	driver.DB.AutoMigrate(
		&stores.User{},
		&stores.UserActionCode{},
		&stores.Client{},
		&stores.Role{},
		&stores.RoleUser{},
		&stores.RoleClient{},
		&stores.RoleUserClient{},
		&stores.ClientAssignment{},
		&stores.UserInvitation{},
		&stores.PermissionRule{},
		&stores.PermissionRuleDetail{},
	)
}
