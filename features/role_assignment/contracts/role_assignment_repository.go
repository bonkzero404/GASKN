package contracts

import (
	"github.com/google/uuid"
)

type RoleAssignmentRepository interface {
	CreateRoleAssignment(roleId uuid.UUID, clientId uuid.UUID, url string, method string) bool
}
