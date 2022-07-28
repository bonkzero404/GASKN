package handlers

import (
	respModel "go-starterkit-project/domain/dto"
	"go-starterkit-project/domain/stores"
	"go-starterkit-project/modules/user/domain/dto"
	"go-starterkit-project/modules/user/domain/interfaces"
	"go-starterkit-project/utils"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	UserService interfaces.UserServiceInterface
}

func NewUserHandler(userService interfaces.UserServiceInterface) *UserHandler {
	return &UserHandler{
		UserService: userService,
	}
}

func (handler *UserHandler) RegisterUser(c *fiber.Ctx) error {
	var request dto.UserCreateRequest

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

	response, err := handler.UserService.CreateUser(c, &request)

	if err != nil {
		re := err.(*respModel.ApiErrorResponse)
		return utils.ApiResponseError(c, re.StatusCode, respModel.Errors{
			Message: utils.Lang(c, "user:err:create:failed"),
			Cause:   err.Error(),
			Inputs:  nil,
		})
	}

	return utils.ApiCreated(c, response)
}

func (handler *UserHandler) UserActivation(c *fiber.Ctx) error {
	var request dto.UserActivationRequest

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

	response, err := handler.UserService.UserActivation(c, request.Email, request.Code)

	if err != nil {
		re := err.(*respModel.ApiErrorResponse)
		return utils.ApiResponseError(c, re.StatusCode, respModel.Errors{
			Message: utils.Lang(c, "user:err:create:activation:failed"),
			Cause:   err.Error(),
			Inputs:  nil,
		})
	}

	return utils.ApiCreated(c, response)
}

func (handler *UserHandler) ReCreateUserActivation(c *fiber.Ctx) error {
	var request dto.UserReActivationRequest

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

	response, err := handler.UserService.CreateUserActivation(c, request.Email, stores.ACTIVATION_CODE)

	if err != nil {
		re := err.(*respModel.ApiErrorResponse)
		return utils.ApiResponseError(c, re.StatusCode, respModel.Errors{
			Message: utils.Lang(c, "user:err:create:re-activation:failed"),
			Cause:   err.Error(),
			Inputs:  nil,
		})
	}

	return utils.ApiCreated(c, response)
}

func (handler *UserHandler) CreateActivationForgotPassword(c *fiber.Ctx) error {
	var request dto.UserForgotPassRequest

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

	response, err := handler.UserService.CreateUserActivation(c, request.Email, stores.FORGOT_PASSWORD)

	if err != nil {
		re := err.(*respModel.ApiErrorResponse)
		return utils.ApiResponseError(c, re.StatusCode, respModel.Errors{
			Message: utils.Lang(c, "user:err:create:forgot-pass:failed"),
			Cause:   err.Error(),
			Inputs:  nil,
		})
	}

	return utils.ApiCreated(c, response)
}

func (handler *UserHandler) UpdatePassword(c *fiber.Ctx) error {
	var request dto.UserForgotPassActRequest

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

	response, err := handler.UserService.UpdatePassword(c, &request)

	if err != nil {
		re := err.(*respModel.ApiErrorResponse)
		return utils.ApiResponseError(c, re.StatusCode, respModel.Errors{
			Message: utils.Lang(c, "user:err:update:password:failed"),
			Cause:   err.Error(),
			Inputs:  nil,
		})
	}

	return utils.ApiCreated(c, response)
}
