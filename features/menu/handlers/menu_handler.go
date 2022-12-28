package handlers

import (
	"github.com/bonkzero404/gaskn/config"
	"github.com/bonkzero404/gaskn/database/stores"
	globalDto "github.com/bonkzero404/gaskn/dto"
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

func (interact *MenuHandler) CreateMenu(c *fiber.Ctx) error {
	var request dto.MenuRequest

	if stat, errRequest := utils.ValidateRequest(c, &request); stat {
		return utils.ApiUnprocessableEntity(c, errRequest)
	}

	response, err := interact.Menu.CreateMenu(c, &request)

	if err != nil {
		re := err.(*globalDto.ApiErrorResponse)
		return utils.ApiResponseError(c, re.StatusCode, globalDto.Errors{
			Message: utils.Lang(c, config.MenuErrCreate),
			Cause:   err.Error(),
			Inputs:  nil,
		})
	}

	return utils.ApiCreated(c, response)
}

func (interact *MenuHandler) GetMenuAll(c *fiber.Ctx) error {
	mode := interact.Menu.ValidationMenuMode(c)
	sort := interact.Menu.ValidationMenuSort(c)
	response, err := interact.Menu.GetMenuAllByType("", mode, sort)

	if err != nil {
		re := err.(*globalDto.ApiErrorResponse)
		return utils.ApiResponseError(c, re.StatusCode, globalDto.Errors{
			Message: utils.Lang(c, config.MenuErrGet),
			Cause:   err.Error(),
			Inputs:  nil,
		})
	}

	return utils.ApiOk(c, response)
}

func (interact *MenuHandler) GetMenuSa(c *fiber.Ctx) error {
	mode := interact.Menu.ValidationMenuMode(c)
	sort := interact.Menu.ValidationMenuSort(c)
	response, err := interact.Menu.GetMenuAllByType(stores.MenuBO, mode, sort)

	if err != nil {
		re := err.(*globalDto.ApiErrorResponse)
		return utils.ApiResponseError(c, re.StatusCode, globalDto.Errors{
			Message: utils.Lang(c, config.MenuErrGet),
			Cause:   err.Error(),
			Inputs:  nil,
		})
	}

	return utils.ApiOk(c, response)
}

func (interact *MenuHandler) GetMenuClient(c *fiber.Ctx) error {
	mode := interact.Menu.ValidationMenuMode(c)
	sort := interact.Menu.ValidationMenuSort(c)
	response, err := interact.Menu.GetMenuAllByType(stores.MenuCL, mode, sort)

	if err != nil {
		re := err.(*globalDto.ApiErrorResponse)
		return utils.ApiResponseError(c, re.StatusCode, globalDto.Errors{
			Message: utils.Lang(c, config.MenuErrGet),
			Cause:   err.Error(),
			Inputs:  nil,
		})
	}

	return utils.ApiOk(c, response)
}
