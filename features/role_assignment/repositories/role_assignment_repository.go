package repositories

import (
	"gaskn/database/driver"
	"gaskn/features/role_assignment/contracts"
	"github.com/casbin/casbin/v2"
	"github.com/google/uuid"
)

type RoleAssignmentRepository struct {
	enforcer *casbin.Enforcer
}

func NewRoleAssignmentRepository(enforcer *casbin.Enforcer) contracts.RoleAssignmentRepository {
	return &RoleAssignmentRepository{
		enforcer: enforcer,
	}
}

func (repository RoleAssignmentRepository) CreateRoleAssignment(roleId uuid.UUID, clientId uuid.UUID, url string, method string) bool {
	repository.enforcer = driver.Enforcer

	if p, _ := repository.enforcer.AddPolicy(
		roleId.String(),
		clientId.String(),
		url,
		method); !p {
		return false
	}

	return true
}
