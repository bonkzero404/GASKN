package interfaces

import (
	"go-starterkit-project/domain/stores"
	"go-starterkit-project/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type RoleClientRepositoryInterface interface {
	CreateRoleClient(role *stores.Role, clientId string) (*stores.Role, error)

	GetRoleClientById(role *stores.RoleClient, id string, clientId string) *gorm.DB

	GetRoleClientByName(role *stores.RoleClient, roleName string, clientId string) *gorm.DB

	GetRoleClientList(role *[]stores.Role, c *fiber.Ctx, clientId string) (*utils.Pagination, error)
}
