package role

import (
	"gaskn/database/driver"
	"gaskn/features/client/handlers"
	"gaskn/features/client/repositories"
	"gaskn/features/client/services"

	"github.com/gofiber/fiber/v2"
)

/*
*
This function is for registering repository - service - handler
*/
func RegisterModule(app *fiber.App) {

	clientRepository := repositories.NewClientRepository(driver.DB)
	clientService := services.NewClientService(clientRepository)
	ClientHandler := handlers.NewClientHandler(clientService)

	routesInit := ApiRoute{
		ClientHandler: *ClientHandler,
	}

	routesInit.Route(app)
}
