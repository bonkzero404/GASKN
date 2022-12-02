package contracts

import (
	"gaskn/features/role_assignment/dto"
	"github.com/gofiber/fiber/v2"
)

type RoleAssignmentService interface {
	CreateRoleAssignment(c *fiber.Ctx, req *dto.RoleAssignmentRequest) (*dto.RoleAssignmentResponse, error)
}
