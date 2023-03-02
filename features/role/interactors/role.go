package interactors

import (
	"github.com/bonkzero404/gaskn/app/utils"
	"github.com/bonkzero404/gaskn/features/role/dto"
)

type Role interface {
	CreateRole(role *dto.RoleRequest) (*dto.RoleResponse, error)

	GetRoleList(page string, limit string, sort string) (*utils.Pagination, error)

	UpdateRole(id string, role *dto.RoleRequest) (*dto.RoleResponse, error)

	DeleteRoleById(id string) (*dto.RoleResponse, error)
}
