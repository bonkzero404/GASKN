package handlers

import (
	"github.com/bonkzero404/gaskn/database/stores"
	globalDto "github.com/bonkzero404/gaskn/dto"
	"github.com/bonkzero404/gaskn/features/user/dto"
	"github.com/bonkzero404/gaskn/features/user/interactors"
	"github.com/bonkzero404/gaskn/utils"

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

	if c.Params("CreateUser") == "create" {
		routeInternal = true
	}

	if stat, errRequest := utils.ValidateRequest(c, &request); stat {
		return utils.ApiUnprocessableEntity(c, errRequest)
	}

	response, err := interact.UserService.CreateUser(c, &request, routeInternal)

	if err != nil {
		re := err.(*globalDto.ApiErrorResponse)
		return utils.ApiResponseError(c, re.StatusCode, globalDto.Errors{
			Message: utils.Lang(c, "user:err:create-failed"),
			Cause:   err.Error(),
			Inputs:  nil,
		})
	}

	return utils.ApiCreated(c, response)
}

func (interact *UserHandler) UserActivation(c *fiber.Ctx) error {
	var request dto.UserActivationRequest

	if stat, errRequest := utils.ValidateRequest(c, &request); stat {
		return utils.ApiUnprocessableEntity(c, errRequest)
	}

	response, err := interact.UserService.UserActivation(c, request.Code)

	if err != nil {
		re := err.(*globalDto.ApiErrorResponse)
		return utils.ApiResponseError(c, re.StatusCode, globalDto.Errors{
			Message: utils.Lang(c, "user:err:activation-failed"),
			Cause:   err.Error(),
			Inputs:  nil,
		})
	}

	return utils.ApiCreated(c, response)
}

func (interact *UserHandler) ReCreateUserActivation(c *fiber.Ctx) error {
	var request dto.UserReActivationRequest

	if stat, errRequest := utils.ValidateRequest(c, &request); stat {
		return utils.ApiUnprocessableEntity(c, errRequest)
	}

	response, err := interact.UserService.CreateUserAction(c, request.Email, stores.ACTIVATION_CODE)

	if err != nil {
		re := err.(*globalDto.ApiErrorResponse)
		return utils.ApiResponseError(c, re.StatusCode, globalDto.Errors{
			Message: utils.Lang(c, "user:err:re-activation-failed"),
			Cause:   err.Error(),
			Inputs:  nil,
		})
	}

	return utils.ApiCreated(c, response)
}

func (interact *UserHandler) CreateActivationForgotPassword(c *fiber.Ctx) error {
	var request dto.UserForgotPassRequest

	if stat, errRequest := utils.ValidateRequest(c, &request); stat {
		return utils.ApiUnprocessableEntity(c, errRequest)
	}

	response, err := interact.UserService.CreateUserAction(c, request.Email, stores.FORGOT_PASSWORD)

	if err != nil {
		re := err.(*globalDto.ApiErrorResponse)
		return utils.ApiResponseError(c, re.StatusCode, globalDto.Errors{
			Message: utils.Lang(c, "user:err:forgot-pass-failed"),
			Cause:   err.Error(),
			Inputs:  nil,
		})
	}

	return utils.ApiCreated(c, response)
}

func (interact *UserHandler) UpdatePassword(c *fiber.Ctx) error {
	var request dto.UserForgotPassActRequest

	if stat, errRequest := utils.ValidateRequest(c, &request); stat {
		return utils.ApiUnprocessableEntity(c, errRequest)
	}

	response, err := interact.UserService.UpdatePassword(c, &request)

	if err != nil {
		re := err.(*globalDto.ApiErrorResponse)
		return utils.ApiResponseError(c, re.StatusCode, globalDto.Errors{
			Message: utils.Lang(c, "user:err:update-pass-failed"),
			Cause:   err.Error(),
			Inputs:  nil,
		})
	}

	return utils.ApiCreated(c, response)
}
