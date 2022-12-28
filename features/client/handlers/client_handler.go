package handlers

import (
	"github.com/bonkzero404/gaskn/config"
	globalDto "github.com/bonkzero404/gaskn/dto"
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

func (interact *ClientHandler) CreateClient(c *fiber.Ctx) error {
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userId := claims["id"].(string)

	var request dto.ClientRequest

	if stat, errRequest := utils.ValidateRequest(c, &request); stat {
		return utils.ApiUnprocessableEntity(c, errRequest)
	}

	response, err := interact.ClientService.CreateClient(c, &request, userId)

	if err != nil {
		re := err.(*globalDto.ApiErrorResponse)
		return utils.ApiResponseError(c, re.StatusCode, globalDto.Errors{
			Message: utils.Lang(c, "client:err:create-failed"),
			Cause:   err.Error(),
			Inputs:  nil,
		})
	}

	return utils.ApiCreated(c, response)
}

func (interact *ClientHandler) UpdateClient(c *fiber.Ctx) error {
	var request dto.ClientRequest

	if stat, errRequest := utils.ValidateRequest(c, &request); stat {
		return utils.ApiUnprocessableEntity(c, errRequest)
	}

	response, err := interact.ClientService.UpdateClient(c, &request)

	if err != nil {
		re := err.(*globalDto.ApiErrorResponse)
		return utils.ApiResponseError(c, re.StatusCode, globalDto.Errors{
			Message: utils.Lang(c, config.ClientErrUpdate),
			Cause:   err.Error(),
			Inputs:  nil,
		})
	}

	return utils.ApiOk(c, response)
}

func (interact *ClientHandler) GetClientByUser(c *fiber.Ctx) error {
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userId := claims["id"].(string)

	response, err := interact.ClientService.GetClientByUser(c, userId)

	if err != nil {
		re := err.(*globalDto.ApiErrorResponse)
		return utils.ApiResponseError(c, re.StatusCode, globalDto.Errors{
			Message: utils.Lang(c, config.ClientErrGet),
			Cause:   err.Error(),
			Inputs:  nil,
		})
	}

	return utils.ApiOk(c, response)
}
