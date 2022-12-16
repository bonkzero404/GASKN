package menu

import (
	"github.com/bonkzero404/gaskn/app/middleware"
	"github.com/bonkzero404/gaskn/features/menu/handlers"
	"github.com/bonkzero404/gaskn/utils"
	"github.com/gofiber/fiber/v2"
)

type ApiRoute struct {
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
		SetRouteDescriptionKeyLang("blabla").
		Execute()

}
