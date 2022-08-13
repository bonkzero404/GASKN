package role

import (
	"go-starterkit-project/database/driver"
	"go-starterkit-project/modules/client/handlers"
	"go-starterkit-project/modules/client/repositories"
	"go-starterkit-project/modules/client/services"

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
