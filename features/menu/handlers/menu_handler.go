package handlers

import (
	"github.com/bonkzero404/gaskn/database/stores"
	responseDto "github.com/bonkzero404/gaskn/dto"
	"github.com/bonkzero404/gaskn/features/menu/dto"
	"github.com/bonkzero404/gaskn/features/menu/interactors"
	"github.com/bonkzero404/gaskn/utils"
	"github.com/gofiber/fiber/v2"
)

type MenuHandler struct {
	Menu interactors.Menu
}

func NewMenuHandler(Menu interactors.Menu) *MenuHandler {
	return &MenuHandler{
		Menu: Menu,
	}
}

func (handler *MenuHandler) CreateMenu(c *fiber.Ctx) error {
	var request dto.MenuRequest

	if err := c.BodyParser(&request); err != nil {
		return utils.ApiUnprocessableEntity(c, responseDto.Errors{
			Message: utils.Lang(c, "global:err:body-parser"),
			Cause:   err.Error(),
			Inputs:  nil,
		})
	}

	errors := utils.ValidateStruct(request, c)
	if errors != nil {
		return utils.ApiErrorValidation(c, responseDto.Errors{
			Message: utils.Lang(c, "global:err:validate"),
			Cause:   utils.Lang(c, "global:err:validate-cause"),
			Inputs:  errors,
		})
	}

	response, err := handler.Menu.CreateMenu(c, &request)

	if err != nil {
		re := err.(*responseDto.ApiErrorResponse)
		return utils.ApiResponseError(c, re.StatusCode, responseDto.Errors{
			Message: utils.Lang(c, "menu:err:create"),
			Cause:   err.Error(),
			Inputs:  nil,
		})
	}

	return utils.ApiCreated(c, response)
}

func (handler *MenuHandler) GetMenuAll(c *fiber.Ctx) error {
	mode := handler.Menu.ValidationMenuMode(c)
	response, err := handler.Menu.GetMenuAllByType("", mode)

	if err != nil {
		re := err.(*responseDto.ApiErrorResponse)
		return utils.ApiResponseError(c, re.StatusCode, responseDto.Errors{
			Message: utils.Lang(c, "menu:err:load"),
			Cause:   err.Error(),
			Inputs:  nil,
		})
	}

	return utils.ApiOk(c, response)
}

func (handler *MenuHandler) GetMenuSa(c *fiber.Ctx) error {
	mode := handler.Menu.ValidationMenuMode(c)
	response, err := handler.Menu.GetMenuAllByType(stores.MenuBO, mode)

	if err != nil {
		re := err.(*responseDto.ApiErrorResponse)
		return utils.ApiResponseError(c, re.StatusCode, responseDto.Errors{
			Message: utils.Lang(c, "menu:err:load"),
			Cause:   err.Error(),
			Inputs:  nil,
		})
	}

	return utils.ApiOk(c, response)
}

func (handler *MenuHandler) GetMenuClient(c *fiber.Ctx) error {
	mode := handler.Menu.ValidationMenuMode(c)
	response, err := handler.Menu.GetMenuAllByType(stores.MenuCL, mode)

	if err != nil {
		re := err.(*responseDto.ApiErrorResponse)
		return utils.ApiResponseError(c, re.StatusCode, responseDto.Errors{
			Message: utils.Lang(c, "menu:err:load"),
			Cause:   err.Error(),
			Inputs:  nil,
		})
	}

	return utils.ApiOk(c, response)
}
