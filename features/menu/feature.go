package menu

import (
	"github.com/bonkzero404/gaskn/driver"
	"github.com/bonkzero404/gaskn/features/menu/handlers"
	interact "github.com/bonkzero404/gaskn/features/menu/interactors/implements"
	repo "github.com/bonkzero404/gaskn/features/menu/repositories/implements"
	"github.com/gofiber/fiber/v2"
)

func RegisterFeature(app *fiber.App) {
	menuRepo := repo.NewMenuRepository(driver.DB)
	menuInteract := interact.NewMenu(menuRepo)
	MenuHandler := handlers.NewMenuHandler(menuInteract)

	routesInit := ApiRoute{
		MenuHandler: *MenuHandler,
	}

	routesInit.Route(app)

}
