package database

import (
	"github.com/bonkzero404/gaskn/database/stores"
	"github.com/bonkzero404/gaskn/infrastructures"
	"log"
)

// MigrateDB /*
func MigrateDB() {
	err := infrastructures.DB.AutoMigrate(
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
		&stores.Menu{},
	)
	if err != nil {
		log.Println(err)
	}
}
