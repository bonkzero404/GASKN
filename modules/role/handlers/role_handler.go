package handlers

import (
	respModel "gaskn/dto"
	"gaskn/modules/role/contracts"
	"gaskn/modules/role/dto"
	"gaskn/utils"

	"github.com/gofiber/fiber/v2"
)

type RoleHandler struct {
	RoleService contracts.RoleService
}

func NewRoleHandler(roleService contracts.RoleService) *RoleHandler {
	return &RoleHandler{
		RoleService: roleService,
	}
}

func (handler *RoleHandler) CreateRole(c *fiber.Ctx) error {
	var request dto.RoleRequest

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

	response, err := handler.RoleService.CreateRole(c, &request)

	if err != nil {
		re := err.(*respModel.ApiErrorResponse)
		return utils.ApiResponseError(c, re.StatusCode, respModel.Errors{
			Message: utils.Lang(c, "role:err:create-failed"),
			Cause:   err.Error(),
			Inputs:  nil,
		})
	}

	return utils.ApiCreated(c, response)
}

func (handler *RoleHandler) GetRoleList(c *fiber.Ctx) error {
	response, err := handler.RoleService.GetRoleList(c)

	if err != nil {
		re := err.(*respModel.ApiErrorResponse)
		return utils.ApiResponseError(c, re.StatusCode, respModel.Errors{
			Message: utils.Lang(c, "role:err:read-failed"),
			Cause:   err.Error(),
			Inputs:  nil,
		})
	}

	return utils.ApiOk(c, response)
}

func (handler *RoleHandler) UpdateRole(c *fiber.Ctx) error {
	var request dto.RoleRequest

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

	roleId := c.Params("id")

	response, err := handler.RoleService.UpdateRole(c, roleId, &request)

	if err != nil {
		re := err.(*respModel.ApiErrorResponse)
		return utils.ApiResponseError(c, re.StatusCode, respModel.Errors{
			Message: utils.Lang(c, "role:err:update-failed"),
			Cause:   err.Error(),
			Inputs:  nil,
		})
	}

	return utils.ApiOk(c, response)
}

func (handler *RoleHandler) DeleteRole(c *fiber.Ctx) error {
	roleId := c.Params("id")

	response, err := handler.RoleService.DeleteRoleById(c, roleId)

	if err != nil {
		re := err.(*respModel.ApiErrorResponse)
		return utils.ApiResponseError(c, re.StatusCode, respModel.Errors{
			Message: utils.Lang(c, "role:err:delete-failed"),
			Cause:   err.Error(),
			Inputs:  nil,
		})
	}

	return utils.ApiOk(c, response)
}
