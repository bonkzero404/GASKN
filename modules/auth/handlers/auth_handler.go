package handlers

import (
	respModel "go-starterkit-project/domain/dto"
	"go-starterkit-project/modules/auth/domain/dto"
	"go-starterkit-project/modules/auth/domain/interfaces"
	"go-starterkit-project/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type AuthHandler struct {
	AuthService interfaces.UserAuthServiceInterface
}

func NewAuthHandler(authService interfaces.UserAuthServiceInterface) *AuthHandler {
	return &AuthHandler{
		AuthService: authService,
	}
}

/**
Authentication handler
*/
func (handler *AuthHandler) Authentication(c *fiber.Ctx) error {
	var request dto.UserAuthRequest

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

	response, err := handler.AuthService.Authenticate(c, &request)

	if err != nil {
		re := err.(*respModel.ApiErrorResponse)
		return utils.ApiResponseError(c, re.StatusCode, respModel.Errors{
			Message: utils.Lang(c, "auth:err:auth-failed"),
			Cause:   err.Error(),
			Inputs:  nil,
		})
	}

	return utils.ApiOk(c, response)
}

/**
Get user profile
*/
func (handler *AuthHandler) GetProfile(c *fiber.Ctx) error {
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	id := claims["id"].(string)

	response, err := handler.AuthService.GetProfile(c, id)

	if err != nil {
		re := err.(*respModel.ApiErrorResponse)
		return utils.ApiResponseError(c, re.StatusCode, respModel.Errors{
			Message: utils.Lang(c, "auth:err:get-profile"),
			Cause:   err.Error(),
			Inputs:  nil,
		})
	}

	return utils.ApiOk(c, response)
}

/**
Refresh token
*/
func (handler *AuthHandler) RefreshToken(c *fiber.Ctx) error {
	token := c.Locals("user").(*jwt.Token)

	response, err := handler.AuthService.RefreshToken(c, token)

	if err != nil {
		re := err.(*respModel.ApiErrorResponse)
		return utils.ApiResponseError(c, re.StatusCode, respModel.Errors{
			Message: utils.Lang(c, "auth:err:get-refresh-token"),
			Cause:   err.Error(),
			Inputs:  nil,
		})
	}

	return utils.ApiOk(c, response)
}
