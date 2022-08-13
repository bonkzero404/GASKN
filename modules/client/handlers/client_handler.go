package handlers

import (
	respModel "go-starterkit-project/domain/dto"
	"go-starterkit-project/modules/client/domain/dto"
	"go-starterkit-project/modules/client/domain/interfaces"
	"go-starterkit-project/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type ClientHandler struct {
	ClientService interfaces.ClientServiceInterface
}

func NewClientHandler(clientService interfaces.ClientServiceInterface) *ClientHandler {
	return &ClientHandler{
		ClientService: clientService,
	}
}

func (handler *ClientHandler) CreateClient(c *fiber.Ctx) error {
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userId := claims["id"].(string)

	var request dto.ClientRequest

	if err := c.BodyParser(&request); err != nil {
		return utils.ApiUnprocessableEntity(c, respModel.Errors{
			Message: utils.Lang(c, "global:err:create:body-parser"),
			Cause:   err.Error(),
			Inputs:  nil,
		})
	}

	errors := utils.ValidateStruct(request, c)
	if errors != nil {
		return utils.ApiErrorValidation(c, respModel.Errors{
			Message: utils.Lang(c, "global:err:create:validate"),
			Cause:   utils.Lang(c, "global:err:create:validate-cause"),
			Inputs:  errors,
		})
	}

	response, err := handler.ClientService.CreateClient(c, &request, userId)

	if err != nil {
		re := err.(*respModel.ApiErrorResponse)
		return utils.ApiResponseError(c, re.StatusCode, respModel.Errors{
			Message: utils.Lang(c, "client:err:create:failed"),
			Cause:   err.Error(),
			Inputs:  nil,
		})
	}

	return utils.ApiCreated(c, response)
}

func (handler *ClientHandler) UpdateClient(c *fiber.Ctx) error {
	var request dto.ClientRequest

	if err := c.BodyParser(&request); err != nil {
		return utils.ApiUnprocessableEntity(c, respModel.Errors{
			Message: utils.Lang(c, "global:err:create:body-parser"),
			Cause:   err.Error(),
			Inputs:  nil,
		})
	}

	errors := utils.ValidateStruct(request, c)
	if errors != nil {
		return utils.ApiErrorValidation(c, respModel.Errors{
			Message: utils.Lang(c, "global:err:create:validate"),
			Cause:   utils.Lang(c, "global:err:create:validate-cause"),
			Inputs:  errors,
		})
	}

	clientId := c.Params("id")

	response, err := handler.ClientService.UpdateClient(c, clientId, &request)

	if err != nil {
		re := err.(*respModel.ApiErrorResponse)
		return utils.ApiResponseError(c, re.StatusCode, respModel.Errors{
			Message: utils.Lang(c, "client:err:update:failed"),
			Cause:   err.Error(),
			Inputs:  nil,
		})
	}

	return utils.ApiOk(c, response)
}
