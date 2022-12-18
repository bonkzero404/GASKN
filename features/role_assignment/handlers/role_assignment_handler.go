package handlers

import (
	responseDto "github.com/bonkzero404/gaskn/dto"
	"github.com/bonkzero404/gaskn/features/role_assignment/dto"
	"github.com/bonkzero404/gaskn/features/role_assignment/interactors"
	"github.com/bonkzero404/gaskn/utils"

	"github.com/gofiber/fiber/v2"
)

type RoleAssignmentHandler struct {
	RoleAssignmentService interactors.RoleAssignment
}

func NewRoleAssignmentHandler(RoleAssignmentService interactors.RoleAssignment) *RoleAssignmentHandler {
	return &RoleAssignmentHandler{
		RoleAssignmentService: RoleAssignmentService,
	}
}

func (service *RoleAssignmentHandler) CreateRoleAssignment(c *fiber.Ctx) error {
	var request dto.RoleAssignmentRequest

	if stat, errRequest := utils.ValidateRequest(c, &request); stat {
		return utils.ApiUnprocessableEntity(c, errRequest)
	}

	response, err := service.RoleAssignmentService.CreateRoleAssignment(c, &request)

	if err != nil {
		re := err.(*responseDto.ApiErrorResponse)
		return utils.ApiResponseError(c, re.StatusCode, responseDto.Errors{
			Message: utils.Lang(c, "client:err:create-failed"),
			Cause:   err.Error(),
			Inputs:  nil,
		})
	}

	return utils.ApiCreated(c, response)
}

func (service *RoleAssignmentHandler) RemoveRoleAssignment(c *fiber.Ctx) error {
	var request dto.RoleAssignmentRequest

	if stat, errRequest := utils.ValidateRequest(c, &request); stat {
		return utils.ApiUnprocessableEntity(c, errRequest)
	}

	response, err := service.RoleAssignmentService.RemoveRoleAssignment(c, &request)

	if err != nil {
		re := err.(*responseDto.ApiErrorResponse)
		return utils.ApiResponseError(c, re.StatusCode, responseDto.Errors{
			Message: utils.Lang(c, "client:err:create-failed"),
			Cause:   err.Error(),
			Inputs:  nil,
		})
	}

	return utils.ApiCreated(c, response)
}

func (service *RoleAssignmentHandler) AssignUserPermitToRole(c *fiber.Ctx) error {
	var request dto.RoleUserAssignment

	if stat, errRequest := utils.ValidateRequest(c, &request); stat {
		return utils.ApiUnprocessableEntity(c, errRequest)
	}

	response, err := service.RoleAssignmentService.AssignUserPermitToRole(c, &request)

	if err != nil {
		re := err.(*responseDto.ApiErrorResponse)
		return utils.ApiResponseError(c, re.StatusCode, responseDto.Errors{
			Message: utils.Lang(c, "client:err:create-failed"),
			Cause:   err.Error(),
			Inputs:  nil,
		})
	}

	return utils.ApiCreated(c, response)
}
