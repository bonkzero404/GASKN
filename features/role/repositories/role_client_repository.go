package repositories

import (
	"github.com/bonkzero404/gaskn/app/utils"
	"github.com/bonkzero404/gaskn/database/stores"
	"github.com/google/uuid"

	"gorm.io/gorm"
)

type RoleClientRepository interface {
	CreateRoleClient(role *stores.Role, clientId string) (*stores.Role, error)

	CreateUserClientRole(userId uuid.UUID, roleId uuid.UUID, clientId uuid.UUID) bool

	GetRoleClientById(role *stores.RoleClient, id string, clientId string) *gorm.DB

	GetRoleClientByName(role *stores.RoleClient, roleName string, clientId string) *gorm.DB

	GetRoleClientList(role *[]stores.Role, clientId string, page int, limit int, sort string) (*utils.Pagination, error)

	GetRoleClientId(role *stores.RoleClient, roleId string, clientId string) *gorm.DB

	GetUserHasClient(clientAssignment *stores.ClientAssignment, userId string, clientId string) *gorm.DB

	GetRoleUser(roleUser *stores.RoleUser, userId string, roleId string) *gorm.DB
}
