package handlers

import (
	"github.com/bonkzero404/gaskn/app/http"
	"github.com/bonkzero404/gaskn/app/http/response"
	"github.com/bonkzero404/gaskn/app/translation"
	"github.com/bonkzero404/gaskn/app/validations"
	"github.com/bonkzero404/gaskn/config"
	"github.com/bonkzero404/gaskn/features/client/dto"
	"github.com/bonkzero404/gaskn/features/client/interactors"
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

	if stat, errRequest := validations.ValidateRequest(c, &request); stat {
		return response.ApiUnprocessableEntity(c, errRequest)
	}

	responseInteract, err := interact.ClientService.CreateClient(c, &request, userId)

	if err != nil {
		re := err.(*http.SetApiErrorResponse)
		return response.ApiResponseError(c, re.StatusCode, http.SetErrors{
			Message: translation.Lang(c, "client:err:create-failed"),
			Cause:   err.Error(),
			Inputs:  nil,
		})
	}

	return response.ApiCreated(c, responseInteract)
}

func (interact *ClientHandler) UpdateClient(c *fiber.Ctx) error {
	var request dto.ClientRequest

	if stat, errRequest := validations.ValidateRequest(c, &request); stat {
		return response.ApiUnprocessableEntity(c, errRequest)
	}

	responseInteract, err := interact.ClientService.UpdateClient(c, &request)

	if err != nil {
		re := err.(*http.SetApiErrorResponse)
		return response.ApiResponseError(c, re.StatusCode, http.SetErrors{
			Message: translation.Lang(c, config.ClientErrUpdate),
			Cause:   err.Error(),
			Inputs:  nil,
		})
	}

	return response.ApiOk(c, responseInteract)
}

func (interact *ClientHandler) GetClientByUser(c *fiber.Ctx) error {
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userId := claims["id"].(string)

	responseInteract, err := interact.ClientService.GetClientByUser(c, userId)

	if err != nil {
		re := err.(*http.SetApiErrorResponse)
		return response.ApiResponseError(c, re.StatusCode, http.SetErrors{
			Message: translation.Lang(c, config.ClientErrGet),
			Cause:   err.Error(),
			Inputs:  nil,
		})
	}

	return response.ApiOk(c, responseInteract)
}
