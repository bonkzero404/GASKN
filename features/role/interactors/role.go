package interactors

import (
	"github.com/bonkzero404/gaskn/app/utils"
	"github.com/bonkzero404/gaskn/features/role/dto"
	"github.com/gofiber/fiber/v2"
)

type Role interface {
	CreateRole(c *fiber.Ctx, role *dto.RoleRequest) (*dto.RoleResponse, error)

	GetRoleList(c *fiber.Ctx) (*utils.Pagination, error)

	UpdateRole(c *fiber.Ctx, id string, role *dto.RoleRequest) (*dto.RoleResponse, error)

	DeleteRoleById(c *fiber.Ctx, id string) (*dto.RoleResponse, error)
}
