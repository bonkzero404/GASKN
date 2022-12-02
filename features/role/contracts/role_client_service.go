package contracts

import (
	"gaskn/features/role/dto"
	"gaskn/utils"

	"github.com/gofiber/fiber/v2"
)

type RoleClientService interface {
	CreateRoleClient(c *fiber.Ctx, role *dto.RoleRequest) (*dto.RoleResponse, error)

	GetRoleClientList(c *fiber.Ctx) (*utils.Pagination, error)

	UpdateRoleClient(c *fiber.Ctx, id string, role *dto.RoleRequest) (*dto.RoleResponse, error)

	DeleteRoleClientById(c *fiber.Ctx, id string) (*dto.RoleResponse, error)
}
