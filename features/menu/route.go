package menu

import (
	"github.com/bonkzero404/gaskn/app/middleware"
	"github.com/bonkzero404/gaskn/config"
	"github.com/bonkzero404/gaskn/features/menu/handlers"
	"github.com/bonkzero404/gaskn/utils"
	"github.com/gofiber/fiber/v2"
)

type ApiRoute struct {
	MenuHandler handlers.MenuHandler
}

type ApiRouteClient struct {
	MenuHandler handlers.MenuHandler
}

// Route /**
func (handler *ApiRoute) Route(app fiber.Router) {
	const endpointGroup string = "/menu"

	var menu = utils.GasknRouter{}

	menu.Set(app).
		Group(utils.SetupApiGroup() + endpointGroup).
		SetGroupName("Menu")

	menu.Post(
		"/",
		middleware.Authenticate(),
		middleware.RateLimiter(5, 30),
		middleware.Permission(),
		handler.MenuHandler.CreateMenu,
	).
		SetRouteName("CreateMenu").
		SetRouteDescriptionKeyLang(config.RouteMenuCreate).
		Execute()

	menu.Get(
		"/",
		middleware.Authenticate(),
		middleware.RateLimiter(5, 30),
		middleware.Permission(),
		handler.MenuHandler.GetMenuAll,
	).
		SetRouteName("GetAllMenu").
		SetRouteDescriptionKeyLang(config.RouteMenuGetAll).
		Execute()

	menu.Get(
		"/sa",
		middleware.Authenticate(),
		middleware.RateLimiter(5, 30),
		middleware.Permission(),
		handler.MenuHandler.GetMenuSa,
	).
		SetRouteName("GetAllMenuSa").
		SetRouteDescriptionKeyLang(config.RouteMenuGetAllSa).
		Execute()

	menu.Get(
		"/client",
		middleware.Authenticate(),
		middleware.RateLimiter(5, 30),
		middleware.Permission(),
		handler.MenuHandler.GetMenuClient,
	).
		SetRouteName("GetAllMenuClient").
		SetRouteDescriptionKeyLang(config.RouteMenuGetAllCl).
		Execute()

}

func (handler *ApiRouteClient) RouteClient(app fiber.Router) {
	const endpointGroup string = "/menu"

	var menuClient = utils.GasknRouter{}

	menuClient.Set(app).
		Group(utils.SetupSubApiGroup() + endpointGroup).
		SetGroupName("Client/Menu")

	menuClient.Get(
		"/",
		middleware.Authenticate(),
		middleware.RateLimiter(5, 30),
		middleware.Permission(),
		handler.MenuHandler.GetMenuClient,
	).
		SetRouteName("GetClientMenu").
		SetRouteDescriptionKeyLang(config.RouteMenuGetAllCl).
		SetRouteTenant(true).
		Execute()
}
