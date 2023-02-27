package interactors

import (
	"github.com/bonkzero404/gaskn/app/utils"
	"github.com/bonkzero404/gaskn/features/role/dto"
	"github.com/gofiber/fiber/v2"
)

type RoleClient interface {
	CreateRoleClient(c *fiber.Ctx, role *dto.RoleRequest) (*dto.RoleResponse, error)

	GetRoleClientList(c *fiber.Ctx) (*utils.Pagination, error)

	UpdateRoleClient(c *fiber.Ctx, id string) (*dto.RoleResponse, error)

	DeleteRoleClientById(c *fiber.Ctx, id string) (*dto.RoleResponse, error)
}
