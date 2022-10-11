package handlers

import (
	respModel "go-starterkit-project/dto"
	"go-starterkit-project/modules/auth/contracts"
	"go-starterkit-project/modules/auth/dto"
	"go-starterkit-project/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type AuthHandler struct {
	AuthService contracts.UserAuthService
}

func NewAuthHandler(authService contracts.UserAuthService) *AuthHandler {
	return &AuthHandler{
		AuthService: authService,
	}
}

/*
*
Authentication handler
*/
func (service *AuthHandler) Authentication(c *fiber.Ctx) error {
	var request dto.UserAuthRequest

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

	response, err := service.AuthService.Authenticate(c, &request)

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

/*
*
Get user profile
*/
func (service *AuthHandler) GetProfile(c *fiber.Ctx) error {
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	id := claims["id"].(string)

	response, err := service.AuthService.GetProfile(c, id)

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

/*
*
Refresh token
*/
func (service *AuthHandler) RefreshToken(c *fiber.Ctx) error {
	token := c.Locals("user").(*jwt.Token)

	response, err := service.AuthService.RefreshToken(c, token)

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
