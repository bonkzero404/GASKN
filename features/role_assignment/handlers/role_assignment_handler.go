package handlers

import (
	"github.com/bonkzero404/gaskn/app/http"
	"github.com/bonkzero404/gaskn/app/http/response"
	"github.com/bonkzero404/gaskn/app/translation"
	"github.com/bonkzero404/gaskn/app/validations"
	"github.com/bonkzero404/gaskn/config"
	"github.com/bonkzero404/gaskn/features/role_assignment/dto"
	"github.com/bonkzero404/gaskn/features/role_assignment/interactors"
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

	if stat, errRequest := validations.ValidateRequest(c, &request); stat {
		return response.ApiUnprocessableEntity(c, errRequest)
	}

	responseInteract, err := interact.RoleAssignmentService.CreateRoleAssignment(c, &request)

	if err != nil {
		re := err.(*http.SetApiErrorResponse)
		return response.ApiResponseError(c, re.StatusCode, http.SetErrors{
			Message: translation.Lang(c, config.RoleAssignErrFailed),
			Cause:   err.Error(),
			Inputs:  nil,
		})
	}

	return response.ApiCreated(c, responseInteract)
}

func (interact *RoleAssignmentHandler) RemoveRoleAssignment(c *fiber.Ctx) error {
	var request dto.RoleAssignmentRequest

	if stat, errRequest := validations.ValidateRequest(c, &request); stat {
		return response.ApiUnprocessableEntity(c, errRequest)
	}

	responseInteract, err := interact.RoleAssignmentService.RemoveRoleAssignment(c, &request)

	if err != nil {
		re := err.(*http.SetApiErrorResponse)
		return response.ApiResponseError(c, re.StatusCode, http.SetErrors{
			Message: translation.Lang(c, config.RoleAssignErrRemovePermit),
			Cause:   err.Error(),
			Inputs:  nil,
		})
	}

	return response.ApiCreated(c, responseInteract)
}

func (interact *RoleAssignmentHandler) AssignUserPermission(c *fiber.Ctx) error {
	var request dto.RoleUserAssignment

	if stat, errRequest := validations.ValidateRequest(c, &request); stat {
		return response.ApiUnprocessableEntity(c, errRequest)
	}

	responseInteract, err := interact.RoleAssignmentService.AssignUserPermission(c, &request)

	if err != nil {
		re := err.(*http.SetApiErrorResponse)
		return response.ApiResponseError(c, re.StatusCode, http.SetErrors{
			Message: translation.Lang(c, config.RoleAssignErrFailed),
			Cause:   err.Error(),
			Inputs:  nil,
		})
	}

	return response.ApiCreated(c, responseInteract)
}

func (interact *RoleAssignmentHandler) GetPermissionByRole(c *fiber.Ctx) error {

	responseInteract, err := interact.RoleAssignmentService.GetPermissionListByRole(c)

	if err != nil {
		re := err.(*http.SetApiErrorResponse)
		return response.ApiResponseError(c, re.StatusCode, http.SetErrors{
			Message: translation.Lang(c, config.RoleAssignErrLoad),
			Cause:   err.Error(),
			Inputs:  nil,
		})
	}

	return response.ApiCreated(c, responseInteract)
}
