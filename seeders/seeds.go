package seeders

import (
	"go-starterkit-project/domain/stores"

	"gorm.io/gorm"
)

func All() []Seed {
	return []Seed{
		{
			Name: "CreateSuperUser",
			Run: func(db *gorm.DB) error {
				return CreateUser(db)
			},
		},
		{
			Name: "CreateRoleSa",
			Run: func(db *gorm.DB) error {
				var roleName string = "Super Administrator"
				var roleDesc string = "User can access all features"
				return CreateRole(db, roleName, roleDesc, stores.SA)
			},
		},
		{
			Name: "CreateRoleOwner",
			Run: func(db *gorm.DB) error {
				var roleName string = "Owner"
				var roleDesc string = "User can access all features from clients"
				return CreateRole(db, roleName, roleDesc, stores.CL)
			},
		},
		{
			Name: "CreateCasbinPermission",
			Run: func(db *gorm.DB) error {
				return CreateCasbinPermission(db)
			},
		},
	}
}
