package handlers

import (
	respModel "gaskn/dto"
	"gaskn/features/role_assignment/contracts"
	"gaskn/features/role_assignment/dto"
	"gaskn/utils"
	"github.com/gofiber/fiber/v2"
)

type RoleAssignmentHandler struct {
	RoleAssignmentService contracts.RoleAssignmentService
}

func NewRoleAssignmentHandler(RoleAssignmentService contracts.RoleAssignmentService) *RoleAssignmentHandler {
	return &RoleAssignmentHandler{
		RoleAssignmentService: RoleAssignmentService,
	}
}

func (service *RoleAssignmentHandler) CreateRoleAssignment(c *fiber.Ctx) error {
	var request dto.RoleAssignmentRequest

	if err := c.BodyParser(&request); err != nil {
		return utils.ApiUnprocessableEntity(c, respModel.Errors{
			Message: utils.Lang(c, "global:err:body-parser"),
			Cause:   err.Error(),
			Inputs:  nil,
		})
	}

	errors := utils.ValidateStruct(request, c)
	if errors != nil {
		return utils.ApiErrorValidation(c, respModel.Errors{
			Message: utils.Lang(c, "global:err:validate"),
			Cause:   utils.Lang(c, "global:err:validate-cause"),
			Inputs:  errors,
		})
	}

	response, err := service.RoleAssignmentService.CreateRoleAssignment(c, &request)

	if err != nil {
		re := err.(*respModel.ApiErrorResponse)
		return utils.ApiResponseError(c, re.StatusCode, respModel.Errors{
			Message: utils.Lang(c, "client:err:create-failed"),
			Cause:   err.Error(),
			Inputs:  nil,
		})
	}

	return utils.ApiCreated(c, response)
}

func (service *RoleAssignmentHandler) RemoveRoleAssignment(c *fiber.Ctx) error {
	var request dto.RoleAssignmentRequest

	if err := c.BodyParser(&request); err != nil {
		return utils.ApiUnprocessableEntity(c, respModel.Errors{
			Message: utils.Lang(c, "global:err:body-parser"),
			Cause:   err.Error(),
			Inputs:  nil,
		})
	}

	errors := utils.ValidateStruct(request, c)
	if errors != nil {
		return utils.ApiErrorValidation(c, respModel.Errors{
			Message: utils.Lang(c, "global:err:validate"),
			Cause:   utils.Lang(c, "global:err:validate-cause"),
			Inputs:  errors,
		})
	}

	response, err := service.RoleAssignmentService.RemoveRoleAssignment(c, &request)

	if err != nil {
		re := err.(*respModel.ApiErrorResponse)
		return utils.ApiResponseError(c, re.StatusCode, respModel.Errors{
			Message: utils.Lang(c, "client:err:create-failed"),
			Cause:   err.Error(),
			Inputs:  nil,
		})
	}

	return utils.ApiCreated(c, response)
}
