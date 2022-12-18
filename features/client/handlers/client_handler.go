package handlers

import (
	responseDto "github.com/bonkzero404/gaskn/dto"
	"github.com/bonkzero404/gaskn/features/client/dto"
	"github.com/bonkzero404/gaskn/features/client/interactors"
	"github.com/bonkzero404/gaskn/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type ClientHandler struct {
	ClientService interactors.Client
}

func NewClientHandler(clientService interactors.Client) *ClientHandler {
	return &ClientHandler{
		ClientService: clientService,
	}
}

func (service *ClientHandler) CreateClient(c *fiber.Ctx) error {
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userId := claims["id"].(string)

	var request dto.ClientRequest

	if stat, errRequest := utils.ValidateRequest(c, &request); stat {
		return utils.ApiUnprocessableEntity(c, errRequest)
	}

	response, err := service.ClientService.CreateClient(c, &request, userId)

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

func (service *ClientHandler) UpdateClient(c *fiber.Ctx) error {
	var request dto.ClientRequest

	if stat, errRequest := utils.ValidateRequest(c, &request); stat {
		return utils.ApiUnprocessableEntity(c, errRequest)
	}

	response, err := service.ClientService.UpdateClient(c, &request)

	if err != nil {
		re := err.(*responseDto.ApiErrorResponse)
		return utils.ApiResponseError(c, re.StatusCode, responseDto.Errors{
			Message: utils.Lang(c, "client:err:update:-failed"),
			Cause:   err.Error(),
			Inputs:  nil,
		})
	}

	return utils.ApiOk(c, response)
}

func (service *ClientHandler) GetClientByUser(c *fiber.Ctx) error {
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userId := claims["id"].(string)

	response, err := service.ClientService.GetClientByUser(c, userId)

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
