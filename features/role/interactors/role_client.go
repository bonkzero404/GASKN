package interactors

import (
	"github.com/bonkzero404/gaskn/features/role/dto"
	"github.com/bonkzero404/gaskn/utils"

	"github.com/gofiber/fiber/v2"
)

type RoleClient interface {
	CreateRoleClient(c *fiber.Ctx, role *dto.RoleRequest) (*dto.RoleResponse, error)

	GetRoleClientList(c *fiber.Ctx) (*utils.Pagination, error)

	UpdateRoleClient(c *fiber.Ctx, id string) (*dto.RoleResponse, error)

	DeleteRoleClientById(c *fiber.Ctx, id string) (*dto.RoleResponse, error)
}
