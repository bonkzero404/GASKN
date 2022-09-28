package interfaces

import (
	"go-starterkit-project/modules/role/domain/dto"
	"go-starterkit-project/utils"

	"github.com/gofiber/fiber/v2"
)

type RoleClientServiceInterface interface {
	CreateRoleClient(c *fiber.Ctx, role *dto.RoleRequest) (*dto.RoleResponse, error)

	GetRoleClientList(c *fiber.Ctx) (*utils.Pagination, error)

	UpdateRoleClient(c *fiber.Ctx, id string, role *dto.RoleRequest) (*dto.RoleResponse, error)

	DeleteRoleClientById(c *fiber.Ctx, id string) (*dto.RoleResponse, error)
}
