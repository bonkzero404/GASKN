package interactors

import (
	"github.com/bonkzero404/gaskn/app/utils"
	"github.com/bonkzero404/gaskn/features/role/dto"
)

type RoleClient interface {
	CreateRoleClient(clientId string, role *dto.RoleRequest) (*dto.RoleResponse, error)

	GetRoleClientList(clientId string, page string, limit string, sort string) (*utils.Pagination, error)

	UpdateRoleClient(clientId string, id string) (*dto.RoleResponse, error)

	DeleteRoleClientById(clientId string, id string) (*dto.RoleResponse, error)
}
