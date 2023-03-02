package handlers

import (
	"github.com/bonkzero404/gaskn/app/http"
	"github.com/bonkzero404/gaskn/app/http/response"
	"github.com/bonkzero404/gaskn/app/translation"
	"github.com/bonkzero404/gaskn/app/validations"
	"github.com/bonkzero404/gaskn/features/role/dto"
	"github.com/bonkzero404/gaskn/features/role/interactors"
	"github.com/gofiber/fiber/v2"
)

type RoleHandler struct {
	Role interactors.Role
}

func NewRoleHandler(roleService interactors.Role) *RoleHandler {
	return &RoleHandler{
		Role: roleService,
	}
}

func (interact *RoleHandler) CreateRole(c *fiber.Ctx) error {
	var request dto.RoleRequest

	if stat, errRequest := validations.ValidateRequest(c, &request); stat {
		return response.ApiUnprocessableEntity(c, errRequest)
	}

	responseInteract, err := interact.Role.CreateRole(&request)

	if err != nil {
		re := err.(*http.SetApiErrorResponse)
		return response.ApiResponseError(c, re.StatusCode, http.SetErrors{
			Message: translation.Lang("role:err:create-failed"),
			Cause:   err.Error(),
			Inputs:  nil,
		})
	}

	return response.ApiCreated(c, responseInteract)
}

func (interact *RoleHandler) GetRoleList(c *fiber.Ctx) error {
	responseInteract, err := interact.Role.GetRoleList(c.Query("page"), c.Query("limit"), c.Query("page"))

	if err != nil {
		re := err.(*http.SetApiErrorResponse)
		return response.ApiResponseError(c, re.StatusCode, http.SetErrors{
			Message: translation.Lang("role:err:read-failed"),
			Cause:   err.Error(),
			Inputs:  nil,
		})
	}

	return response.ApiOk(c, responseInteract)
}

func (interact *RoleHandler) UpdateRole(c *fiber.Ctx) error {
	var request dto.RoleRequest

	if stat, errRequest := validations.ValidateRequest(c, &request); stat {
		return response.ApiUnprocessableEntity(c, errRequest)
	}

	roleId := c.Params("id")

	responseInteract, err := interact.Role.UpdateRole(roleId, &request)

	if err != nil {
		re := err.(*http.SetApiErrorResponse)
		return response.ApiResponseError(c, re.StatusCode, http.SetErrors{
			Message: translation.Lang("role:err:update-failed"),
			Cause:   err.Error(),
			Inputs:  nil,
		})
	}

	return response.ApiOk(c, responseInteract)
}

func (interact *RoleHandler) DeleteRole(c *fiber.Ctx) error {
	roleId := c.Params("id")

	responseInteract, err := interact.Role.DeleteRoleById(roleId)

	if err != nil {
		re := err.(*http.SetApiErrorResponse)
		return response.ApiResponseError(c, re.StatusCode, http.SetErrors{
			Message: translation.Lang("role:err:delete-failed"),
			Cause:   err.Error(),
			Inputs:  nil,
		})
	}

	return response.ApiOk(c, responseInteract)
}
