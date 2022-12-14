package repositories

import (
	"github.com/bonkzero404/gaskn/database/stores"
	"github.com/bonkzero404/gaskn/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type RoleRepository interface {
	CreateRole(role *stores.Role) *gorm.DB

	UpdateRoleById(role *stores.Role) *gorm.DB

	DeleteRoleById(role *stores.Role) *gorm.DB

	GetRoleById(role *stores.Role, id string) *gorm.DB

	GetRoleList(role *[]stores.Role, c *fiber.Ctx) (*utils.Pagination, error)
}
