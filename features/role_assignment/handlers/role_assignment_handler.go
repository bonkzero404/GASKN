package handlers

import (
	"github.com/bonkzero404/gaskn/config"
	globalDto "github.com/bonkzero404/gaskn/dto"
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

func (interact *RoleAssignmentHandler) CreateRoleAssignment(c *fiber.Ctx) error {
	var request dto.RoleAssignmentRequest

	if stat, errRequest := utils.ValidateRequest(c, &request); stat {
		return utils.ApiUnprocessableEntity(c, errRequest)
	}

	response, err := interact.RoleAssignmentService.CreateRoleAssignment(c, &request)

	if err != nil {
		re := err.(*globalDto.ApiErrorResponse)
		return utils.ApiResponseError(c, re.StatusCode, globalDto.Errors{
			Message: utils.Lang(c, config.RoleAssignErrFailed),
			Cause:   err.Error(),
			Inputs:  nil,
		})
	}

	return utils.ApiCreated(c, response)
}

func (interact *RoleAssignmentHandler) RemoveRoleAssignment(c *fiber.Ctx) error {
	var request dto.RoleAssignmentRequest

	if stat, errRequest := utils.ValidateRequest(c, &request); stat {
		return utils.ApiUnprocessableEntity(c, errRequest)
	}

	response, err := interact.RoleAssignmentService.RemoveRoleAssignment(c, &request)

	if err != nil {
		re := err.(*globalDto.ApiErrorResponse)
		return utils.ApiResponseError(c, re.StatusCode, globalDto.Errors{
			Message: utils.Lang(c, config.RoleAssignErrRemovePermit),
			Cause:   err.Error(),
			Inputs:  nil,
		})
	}

	return utils.ApiCreated(c, response)
}

func (interact *RoleAssignmentHandler) AssignUserPermission(c *fiber.Ctx) error {
	var request dto.RoleUserAssignment

	if stat, errRequest := utils.ValidateRequest(c, &request); stat {
		return utils.ApiUnprocessableEntity(c, errRequest)
	}

	response, err := interact.RoleAssignmentService.AssignUserPermission(c, &request)

	if err != nil {
		re := err.(*globalDto.ApiErrorResponse)
		return utils.ApiResponseError(c, re.StatusCode, globalDto.Errors{
			Message: utils.Lang(c, config.RoleAssignErrFailed),
			Cause:   err.Error(),
			Inputs:  nil,
		})
	}

	return utils.ApiCreated(c, response)
}

func (interact *RoleAssignmentHandler) GetPermissionByRole(c *fiber.Ctx) error {

	response, err := interact.RoleAssignmentService.GetPermissionListByRole(c)

	if err != nil {
		re := err.(*globalDto.ApiErrorResponse)
		return utils.ApiResponseError(c, re.StatusCode, globalDto.Errors{
			Message: utils.Lang(c, config.RoleAssignErrLoad),
			Cause:   err.Error(),
			Inputs:  nil,
		})
	}

	return utils.ApiCreated(c, response)
}
