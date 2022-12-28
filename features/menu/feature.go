package menu

import (
	"github.com/bonkzero404/gaskn/config"
	"github.com/bonkzero404/gaskn/driver"
	"github.com/bonkzero404/gaskn/features/menu/handlers"
	menuInteract "github.com/bonkzero404/gaskn/features/menu/interactors/implements"
	menuRepo "github.com/bonkzero404/gaskn/features/menu/repositories/implements"
	"github.com/gofiber/fiber/v2"
)

func RegisterFeature(app *fiber.App) {
	menuRepo := menuRepo.NewMenuRepository(driver.DB)
	menu := menuInteract.NewMenu(menuRepo)
	MenuHandler := handlers.NewMenuHandler(menu)

	var routesInitTenant = ApiRouteClient{}

	routesInit := ApiRoute{
		MenuHandler: *MenuHandler,
	}

	routesInit.Route(app)

	/////////////////////////
	// If tenant is enabled
	/////////////////////////
	if config.Config("TENANCY") == "true" {
		routesInitTenant = ApiRouteClient{
			MenuHandler: *MenuHandler,
		}
		routesInitTenant.RouteClient(app)
	}

}
