package handlers

import (
	responseDto "github.com/bonkzero404/gaskn/dto"
	"github.com/bonkzero404/gaskn/features/role/dto"
	"github.com/bonkzero404/gaskn/features/role/interactors"
	"github.com/bonkzero404/gaskn/utils"

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

func (service *RoleClientHandler) CreateClientRole(c *fiber.Ctx) error {
	var request dto.RoleRequest

	if stat, errRequest := utils.ValidateRequest(c, &request); stat {
		return utils.ApiUnprocessableEntity(c, errRequest)
	}

	response, err := service.RoleClientService.CreateRoleClient(c, &request)

	if err != nil {
		re := err.(*responseDto.ApiErrorResponse)
		return utils.ApiResponseError(c, re.StatusCode, responseDto.Errors{
			Message: utils.Lang(c, "role:err:create-failed"),
			Cause:   err.Error(),
			Inputs:  nil,
		})
	}

	return utils.ApiCreated(c, response)
}

func (service *RoleClientHandler) GetRoleClientList(c *fiber.Ctx) error {
	response, err := service.RoleClientService.GetRoleClientList(c)

	if err != nil {
		re := err.(*responseDto.ApiErrorResponse)
		return utils.ApiResponseError(c, re.StatusCode, responseDto.Errors{
			Message: utils.Lang(c, "role:err:read-failed"),
			Cause:   err.Error(),
			Inputs:  nil,
		})
	}

	return utils.ApiOk(c, response)
}

func (service *RoleClientHandler) UpdateRoleClient(c *fiber.Ctx) error {
	var request dto.RoleRequest

	if stat, errRequest := utils.ValidateRequest(c, &request); stat {
		return utils.ApiUnprocessableEntity(c, errRequest)
	}

	roleId := c.Params("id")

	response, err := service.RoleClientService.UpdateRoleClient(c, roleId)

	if err != nil {
		re := err.(*responseDto.ApiErrorResponse)
		return utils.ApiResponseError(c, re.StatusCode, responseDto.Errors{
			Message: utils.Lang(c, "role:err:update-failed"),
			Cause:   err.Error(),
			Inputs:  nil,
		})
	}

	return utils.ApiOk(c, response)
}

func (service *RoleClientHandler) DeleteRoleClient(c *fiber.Ctx) error {
	roleId := c.Params("id")

	response, err := service.RoleClientService.DeleteRoleClientById(c, roleId)

	if err != nil {
		re := err.(*responseDto.ApiErrorResponse)
		return utils.ApiResponseError(c, re.StatusCode, responseDto.Errors{
			Message: utils.Lang(c, "role:err:delete-failed"),
			Cause:   err.Error(),
			Inputs:  nil,
		})
	}

	return utils.ApiOk(c, response)
}
