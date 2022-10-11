package database

import (
	"gaskn/database/driver"
	"gaskn/database/stores"
)

/*
*
This function is used for auto migration and is loaded
into the main function
*/
func MigrateDB() {
	driver.DB.AutoMigrate(
		&stores.User{},
		&stores.UserActivation{},
		&stores.Client{},
		&stores.Role{},
		&stores.RoleUser{},
		&stores.RoleClient{},
		&stores.RoleUserClient{},
		&stores.ClientAssignment{},
	)
}
