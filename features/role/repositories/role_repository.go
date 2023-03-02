package repositories

import (
	"github.com/bonkzero404/gaskn/app/utils"
	"github.com/bonkzero404/gaskn/database/stores"
	"gorm.io/gorm"
)

type RoleRepository interface {
	CreateRole(role *stores.Role) *gorm.DB

	UpdateRoleById(role *stores.Role) *gorm.DB

	DeleteRoleById(role *stores.Role) *gorm.DB

	GetRoleById(role *stores.Role, id string) *gorm.DB

	GetRoleList(role *[]stores.Role, page int, limit int, sort string) (*utils.Pagination, error)
}
