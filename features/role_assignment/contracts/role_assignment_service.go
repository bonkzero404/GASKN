package contracts

import (
	"gaskn/database/stores"
	"gaskn/features/role_assignment/dto"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type RoleAssignmentService interface {
	CheckExistsRoleAssignment(c *fiber.Ctx, clientIdUuid uuid.UUID, roleIdUuid uuid.UUID) (*stores.RoleClient, error)
	CreateRoleAssignment(c *fiber.Ctx, req *dto.RoleAssignmentRequest) (*dto.RoleAssignmentResponse, error)
	RemoveRoleAssignment(c *fiber.Ctx, req *dto.RoleAssignmentRequest) (*dto.RoleAssignmentResponse, error)
}
