package database

import (
	"gaskn/database/stores"
	"gaskn/driver"
	"log"
)

// MigrateDB /*
func MigrateDB() {
	err := driver.DB.AutoMigrate(
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
	if err != nil {
		log.Println(err)
	}
}
