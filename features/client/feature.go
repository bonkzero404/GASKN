package role

import (
	"gaskn/driver"
	"gaskn/features/client/handlers"
	"gaskn/features/client/interactors/implements"
	implements2 "gaskn/features/client/repositories/implements"
	userRepo "gaskn/features/user/repositories/implements"

	"github.com/gofiber/fiber/v2"
)

// RegisterFeature /*
func RegisterFeature(app *fiber.App) {
	userReposiry := userRepo.NewUserRepository(driver.DB)
	clientRepository := implements2.NewClientRepository(driver.DB)
	clientService := implements.NewClient(clientRepository, userReposiry)
	ClientHandler := handlers.NewClientHandler(clientService)

	routesInit := ApiRoute{
		ClientHandler: *ClientHandler,
	}

	routesInit.Route(app)
}
