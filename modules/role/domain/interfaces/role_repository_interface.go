package interfaces

import (
	"go-starterkit-project/domain/stores"
	"go-starterkit-project/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type RoleRepositoryInterface interface {
	CreateRole(role *stores.Role) *gorm.DB

	UpdateRoleById(role *stores.Role) *gorm.DB

	DeleteRoleById(role *stores.Role) *gorm.DB

	GetRoleById(role *stores.Role, id string) *gorm.DB

	GetRoleList(role *[]stores.Role, c *fiber.Ctx) (*utils.Pagination, error)

	GetClientsByUser(roleClient *[]stores.RoleUser, userId string) *gorm.DB
}
