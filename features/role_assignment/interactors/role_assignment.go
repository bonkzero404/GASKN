package interactors

import (
	"github.com/bonkzero404/gaskn/database/stores"
	"github.com/bonkzero404/gaskn/features/role_assignment/dto"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type RoleAssignment interface {
	CheckExistsRoleAssignment(c *fiber.Ctx, clientIdUuid uuid.UUID, roleIdUuid uuid.UUID) (*stores.RoleClient, error)

	CheckExistsRoleUserAssignment(c *fiber.Ctx, userId uuid.UUID, clientIdUuid uuid.UUID) (*stores.ClientAssignment, error)

	CreateRoleAssignment(c *fiber.Ctx, req *dto.RoleAssignmentRequest) (*dto.RoleAssignmentResponse, error)

	RemoveRoleAssignment(c *fiber.Ctx, req *dto.RoleAssignmentRequest) (*dto.RoleAssignmentResponse, error)

	AssignUserPermitToRole(c *fiber.Ctx, req *dto.RoleUserAssignment) (*dto.RoleAssignmentResponse, error)
}
