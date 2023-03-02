package handlers

import (
	"github.com/bonkzero404/gaskn/app/http"
	"github.com/bonkzero404/gaskn/app/http/response"
	"github.com/bonkzero404/gaskn/app/translation"
	"github.com/bonkzero404/gaskn/app/validations"
	"github.com/bonkzero404/gaskn/config"
	"github.com/bonkzero404/gaskn/database/stores"
	"github.com/bonkzero404/gaskn/features/user/dto"
	"github.com/bonkzero404/gaskn/features/user/interactors"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	UserService interactors.User
}

func NewUserHandler(userService interactors.User) *UserHandler {
	return &UserHandler{
		UserService: userService,
	}
}

func (interact *UserHandler) CreateUser(c *fiber.Ctx) error {
	var request dto.UserCreateRequest
	var routeInternal = false
	var clientId = c.Params(config.Config("API_CLIENT_PARAM"))

	if c.Params("CreateUser") == "create" {
		routeInternal = true
	}

	if stat, errRequest := validations.ValidateRequest(c, &request); stat {
		return response.ApiUnprocessableEntity(c, errRequest)
	}

	responseInteract, err := interact.UserService.CreateUser(clientId, &request, routeInternal)

	if err != nil {
		re := err.(*http.SetApiErrorResponse)
		return response.ApiResponseError(c, re.StatusCode, http.SetErrors{
			Message: translation.Lang("user:err:create-failed"),
			Cause:   err.Error(),
			Inputs:  nil,
		})
	}

	return response.ApiCreated(c, responseInteract)
}

func (interact *UserHandler) UserActivation(c *fiber.Ctx) error {
	var request dto.UserActivationRequest

	if stat, errRequest := validations.ValidateRequest(c, &request); stat {
		return response.ApiUnprocessableEntity(c, errRequest)
	}

	responseInteract, err := interact.UserService.UserActivation(request.Code)

	if err != nil {
		re := err.(*http.SetApiErrorResponse)
		return response.ApiResponseError(c, re.StatusCode, http.SetErrors{
			Message: translation.Lang("user:err:activation-failed"),
			Cause:   err.Error(),
			Inputs:  nil,
		})
	}

	return response.ApiCreated(c, responseInteract)
}

func (interact *UserHandler) ReCreateUserActivation(c *fiber.Ctx) error {
	var request dto.UserReActivationRequest

	if stat, errRequest := validations.ValidateRequest(c, &request); stat {
		return response.ApiUnprocessableEntity(c, errRequest)
	}

	responseInteract, err := interact.UserService.CreateUserAction(request.Email, stores.ACTIVATION_CODE)

	if err != nil {
		re := err.(*http.SetApiErrorResponse)
		return response.ApiResponseError(c, re.StatusCode, http.SetErrors{
			Message: translation.Lang("user:err:re-activation-failed"),
			Cause:   err.Error(),
			Inputs:  nil,
		})
	}

	return response.ApiCreated(c, responseInteract)
}

func (interact *UserHandler) CreateActivationForgotPassword(c *fiber.Ctx) error {
	var request dto.UserForgotPassRequest

	if stat, errRequest := validations.ValidateRequest(c, &request); stat {
		return response.ApiUnprocessableEntity(c, errRequest)
	}

	responseInteract, err := interact.UserService.CreateUserAction(request.Email, stores.FORGOT_PASSWORD)

	if err != nil {
		re := err.(*http.SetApiErrorResponse)
		return response.ApiResponseError(c, re.StatusCode, http.SetErrors{
			Message: translation.Lang("user:err:forgot-pass-failed"),
			Cause:   err.Error(),
			Inputs:  nil,
		})
	}

	return response.ApiCreated(c, responseInteract)
}

func (interact *UserHandler) UpdatePassword(c *fiber.Ctx) error {
	var request dto.UserForgotPassActRequest

	if stat, errRequest := validations.ValidateRequest(c, &request); stat {
		return response.ApiUnprocessableEntity(c, errRequest)
	}

	responseInteract, err := interact.UserService.UpdatePassword(&request)

	if err != nil {
		re := err.(*http.SetApiErrorResponse)
		return response.ApiResponseError(c, re.StatusCode, http.SetErrors{
			Message: translation.Lang("user:err:update-pass-failed"),
			Cause:   err.Error(),
			Inputs:  nil,
		})
	}

	return response.ApiCreated(c, responseInteract)
}
