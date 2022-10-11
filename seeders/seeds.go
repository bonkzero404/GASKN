package seeders

import (
	"go-starterkit-project/config"
	"go-starterkit-project/database/stores"

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
				var roleName string = config.Config("ADMIN_ROLENAME")
				var roleDesc string = "User can access all features"
				return CreateRole(db, roleName, roleDesc, stores.SA)
			},
		},
		{
			Name: "CreateRoleOwner",
			Run: func(db *gorm.DB) error {
				var roleName string = config.Config("CLIENT_ROLE_OWNER_NAME")
				var roleDesc string = "User can access all features from clients"
				return CreateRole(db, roleName, roleDesc, stores.CL)
			},
		},
		{
			Name: "CreateRoleUser",
			Run: func(db *gorm.DB) error {
				return CreateUserRole(db)
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
