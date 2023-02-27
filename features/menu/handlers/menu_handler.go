package handlers

import (
	"github.com/bonkzero404/gaskn/app/http"
	"github.com/bonkzero404/gaskn/app/http/response"
	"github.com/bonkzero404/gaskn/app/translation"
	"github.com/bonkzero404/gaskn/app/validations"
	"github.com/bonkzero404/gaskn/config"
	"github.com/bonkzero404/gaskn/database/stores"
	"github.com/bonkzero404/gaskn/features/menu/dto"
	"github.com/bonkzero404/gaskn/features/menu/interactors"
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

	if stat, errRequest := validations.ValidateRequest(c, &request); stat {
		return response.ApiUnprocessableEntity(c, errRequest)
	}

	responseInteract, err := interact.Menu.CreateMenu(c, &request)

	if err != nil {
		re := err.(*http.SetApiErrorResponse)
		return response.ApiResponseError(c, re.StatusCode, http.SetErrors{
			Message: translation.Lang(c, config.MenuErrCreate),
			Cause:   err.Error(),
			Inputs:  nil,
		})
	}

	return response.ApiCreated(c, responseInteract)
}

func (interact *MenuHandler) GetMenuAll(c *fiber.Ctx) error {
	mode := interact.Menu.ValidationMenuMode(c)
	sort := interact.Menu.ValidationMenuSort(c)
	responseInteract, err := interact.Menu.GetMenuAllByType(c, "", mode, sort)

	if err != nil {
		re := err.(*http.SetApiErrorResponse)
		return response.ApiResponseError(c, re.StatusCode, http.SetErrors{
			Message: translation.Lang(c, config.MenuErrGet),
			Cause:   err.Error(),
			Inputs:  nil,
		})
	}

	return response.ApiOk(c, responseInteract)
}

func (interact *MenuHandler) GetMenuSa(c *fiber.Ctx) error {
	mode := interact.Menu.ValidationMenuMode(c)
	sort := interact.Menu.ValidationMenuSort(c)
	responseInteract, err := interact.Menu.GetMenuAllByType(c, stores.MenuBO, mode, sort)

	if err != nil {
		re := err.(*http.SetApiErrorResponse)
		return response.ApiResponseError(c, re.StatusCode, http.SetErrors{
			Message: translation.Lang(c, config.MenuErrGet),
			Cause:   err.Error(),
			Inputs:  nil,
		})
	}

	return response.ApiOk(c, responseInteract)
}

func (interact *MenuHandler) GetMenuClient(c *fiber.Ctx) error {
	mode := interact.Menu.ValidationMenuMode(c)
	sort := interact.Menu.ValidationMenuSort(c)
	responseInteract, err := interact.Menu.GetMenuAllByType(c, stores.MenuCL, mode, sort)

	if err != nil {
		re := err.(*http.SetApiErrorResponse)
		return response.ApiResponseError(c, re.StatusCode, http.SetErrors{
			Message: translation.Lang(c, config.MenuErrGet),
			Cause:   err.Error(),
			Inputs:  nil,
		})
	}

	return response.ApiOk(c, responseInteract)
}
