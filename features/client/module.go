package role

import (
	"gaskn/database/driver"
	"gaskn/features/client/handlers"
	"gaskn/features/client/repositories"
	"gaskn/features/client/services"
	userRepo "gaskn/features/user/repositories"

	"github.com/gofiber/fiber/v2"
)

/*
*
This function is for registering repository - service - handler
*/
func RegisterModule(app *fiber.App) {
	userReposiry := userRepo.NewUserRepository(driver.DB)
	clientRepository := repositories.NewClientRepository(driver.DB)
	clientService := services.NewClientService(clientRepository, userReposiry)
	ClientHandler := handlers.NewClientHandler(clientService)

	routesInit := ApiRoute{
		ClientHandler: *ClientHandler,
	}

	routesInit.Route(app)
}
