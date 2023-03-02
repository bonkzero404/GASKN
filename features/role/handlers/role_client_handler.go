package handlers

import (
	"github.com/bonkzero404/gaskn/app/http"
	"github.com/bonkzero404/gaskn/app/http/response"
	"github.com/bonkzero404/gaskn/app/translation"
	"github.com/bonkzero404/gaskn/app/validations"
	"github.com/bonkzero404/gaskn/config"
	"github.com/bonkzero404/gaskn/features/role/dto"
	"github.com/bonkzero404/gaskn/features/role/interactors"
	"github.com/gofiber/fiber/v2"
)

type RoleClientHandler struct {
	RoleClientService interactors.RoleClient
}

func NewRoleClientHandler(roleClientService interactors.RoleClient) *RoleClientHandler {
	return &RoleClientHandler{
		RoleClientService: roleClientService,
	}
}

func (interact *RoleClientHandler) CreateClientRole(c *fiber.Ctx) error {
	var request dto.RoleRequest
	var clientId = c.Params(config.Config("API_CLIENT_PARAM"))

	if stat, errRequest := validations.ValidateRequest(c, &request); stat {
		return response.ApiUnprocessableEntity(c, errRequest)
	}

	responseInteract, err := interact.RoleClientService.CreateRoleClient(clientId, &request)

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

func (interact *RoleClientHandler) GetRoleClientList(c *fiber.Ctx) error {
	var clientId = c.Params(config.Config("API_CLIENT_PARAM"))
	responseInteract, err := interact.RoleClientService.GetRoleClientList(clientId, c.Query("page"), c.Query("limit"), c.Query("sort"))

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

func (interact *RoleClientHandler) UpdateRoleClient(c *fiber.Ctx) error {
	var request dto.RoleRequest
	var clientId = c.Params(config.Config("API_CLIENT_PARAM"))

	if stat, errRequest := validations.ValidateRequest(c, &request); stat {
		return response.ApiUnprocessableEntity(c, errRequest)
	}

	roleId := c.Params("id")

	responseInteract, err := interact.RoleClientService.UpdateRoleClient(clientId, roleId)

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

func (interact *RoleClientHandler) DeleteRoleClient(c *fiber.Ctx) error {
	roleId := c.Params("id")
	var clientId = c.Params(config.Config("API_CLIENT_PARAM"))

	responseInteract, err := interact.RoleClientService.DeleteRoleClientById(clientId, roleId)

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
