package contracts

import (
	"gaskn/database/stores"
	"gaskn/utils"
	"github.com/google/uuid"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type RoleClientRepository interface {
	CreateRoleClient(role *stores.Role, clientId string) (*stores.Role, error)

	CreateUserClientRole(userId uuid.UUID, roleId uuid.UUID, clientId uuid.UUID) bool

	GetRoleClientById(role *stores.RoleClient, id string, clientId string) *gorm.DB

	GetRoleClientByName(role *stores.RoleClient, roleName string, clientId string) *gorm.DB

	GetRoleClientList(role *[]stores.Role, c *fiber.Ctx, clientId string) (*utils.Pagination, error)

	GetRoleClientId(role *stores.RoleClient, roleId string, clientId string) *gorm.DB

	GetUserHasClient(clientAssignment *stores.ClientAssignment, userId string, clientId string) *gorm.DB
}
