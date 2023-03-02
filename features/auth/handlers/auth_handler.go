package handlers

import (
	"fmt"
	"github.com/bonkzero404/gaskn/app/http"
	"github.com/bonkzero404/gaskn/app/http/response"
	"github.com/bonkzero404/gaskn/app/translation"
	"github.com/bonkzero404/gaskn/app/validations"
	"github.com/bonkzero404/gaskn/config"
	"github.com/bonkzero404/gaskn/features/auth/dto"
	"github.com/bonkzero404/gaskn/features/auth/interactors"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type AuthHandler struct {
	AuthService interactors.UserAuth
}

func NewAuthHandler(authService interactors.UserAuth) *AuthHandler {
	return &AuthHandler{
		AuthService: authService,
	}
}

// Authentication /*
func (interact *AuthHandler) Authentication(c *fiber.Ctx) error {
	var request dto.UserAuthRequest

	fmt.Println(translation.LangContext)

	if stat, errRequest := validations.ValidateRequest(c, &request); stat {
		return response.ApiUnprocessableEntity(c, errRequest)
	}

	responseInteract, err := interact.AuthService.Authenticate(&request)

	if err != nil {
		re := err.(*http.SetApiErrorResponse)
		return response.ApiResponseError(c, re.StatusCode, http.SetErrors{
			Message: translation.Lang(config.GlobalErrAuthFailed),
			Cause:   err.Error(),
			Inputs:  nil,
		})
	}

	return response.ApiOk(c, responseInteract)
}

// GetProfile /*
func (interact *AuthHandler) GetProfile(c *fiber.Ctx) error {
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	id := claims["id"].(string)

	responseInteract, err := interact.AuthService.GetProfile(id)

	if err != nil {
		re := err.(*http.SetApiErrorResponse)
		return response.ApiResponseError(c, re.StatusCode, http.SetErrors{
			Message: translation.Lang(config.GlobalErrUnknown),
			Cause:   err.Error(),
			Inputs:  nil,
		})
	}

	return response.ApiOk(c, responseInteract)
}

// RefreshToken /*
func (interact *AuthHandler) RefreshToken(c *fiber.Ctx) error {
	token := c.Locals("user").(*jwt.Token)

	responseInteract, err := interact.AuthService.RefreshToken(token)

	if err != nil {
		re := err.(*http.SetApiErrorResponse)
		return response.ApiResponseError(c, re.StatusCode, http.SetErrors{
			Message: translation.Lang(config.AuthErruserNotActive),
			Cause:   err.Error(),
			Inputs:  nil,
		})
	}

	return response.ApiOk(c, responseInteract)
}
