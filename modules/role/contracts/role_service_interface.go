package contracts

import (
	"go-starterkit-project/modules/role/dto"
	"go-starterkit-project/utils"

	"github.com/gofiber/fiber/v2"
)

type RoleServiceInterface interface {
	CreateRole(c *fiber.Ctx, role *dto.RoleRequest) (*dto.RoleResponse, error)

	GetRoleList(c *fiber.Ctx) (*utils.Pagination, error)

	UpdateRole(c *fiber.Ctx, id string, role *dto.RoleRequest) (*dto.RoleResponse, error)

	DeleteRoleById(c *fiber.Ctx, id string) (*dto.RoleResponse, error)
}
