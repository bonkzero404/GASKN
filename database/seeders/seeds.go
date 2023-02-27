package seeders

import (
	"github.com/bonkzero404/gaskn/config"
	"github.com/bonkzero404/gaskn/database/stores"

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
				var roleName = config.Config("ADMIN_ROLENAME")
				var roleDesc = "User can access all features"
				return CreateRole(db, roleName, roleDesc, stores.SA)
			},
		},
		{
			Name: "CreateRoleOwner",
			Run: func(db *gorm.DB) error {
				var roleName = config.Config("CLIENT_ROLE_OWNER_NAME")
				var roleDesc = "User can access all features from clients"
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
