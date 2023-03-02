package interactors

import (
	"github.com/bonkzero404/gaskn/database/stores"
	"github.com/bonkzero404/gaskn/features/role_assignment/dto"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type RoleAssignment interface {
	CheckExistsRoleAssignment(clientIdUuid uuid.UUID, roleIdUuid uuid.UUID) (*stores.RoleClient, error)

	CheckExistsRoleUserAssignment(userId uuid.UUID, clientIdUuid uuid.UUID) (*stores.ClientAssignment, error)

	CreateRoleAssignment(clientId string, req *dto.RoleAssignmentRequest) (*dto.RoleAssignmentResponse, error)

	RemoveRoleAssignment(clientId string, req *dto.RoleAssignmentRequest) (*dto.RoleAssignmentResponse, error)

	AssignUserPermission(clientId string, req *dto.RoleUserAssignment) (*dto.RoleAssignmentResponse, error)

	GetPermissionListByRole(c *fiber.Ctx) (*[]dto.RoleAssignmentListResponse, error)
}
