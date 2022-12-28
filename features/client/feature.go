package role

import (
	"github.com/bonkzero404/gaskn/config"
	"github.com/bonkzero404/gaskn/driver"
	"github.com/bonkzero404/gaskn/features/client/handlers"
	clientInteract "github.com/bonkzero404/gaskn/features/client/interactors/implements"
	clientRepo "github.com/bonkzero404/gaskn/features/client/repositories/implements"
	userRepo "github.com/bonkzero404/gaskn/features/user/repositories/implements"

	"github.com/gofiber/fiber/v2"
)

// RegisterFeature /*
func RegisterFeature(app *fiber.App) {
	userReposiry := userRepo.NewUserRepository(driver.DB)
	clientRepository := clientRepo.NewClientRepository(driver.DB)
	client := clientInteract.NewClient(clientRepository, userReposiry)
	ClientHandler := handlers.NewClientHandler(client)

	var routesInitTenant = ApiRouteClient{}

	routesInit := ApiRoute{
		ClientHandler: *ClientHandler,
	}

	routesInit.Route(app)

	/////////////////////////
	// If tenant is enabled
	/////////////////////////
	if config.Config("TENANCY") == "true" {
		routesInitTenant = ApiRouteClient{
			ClientHandler: *ClientHandler,
		}
		routesInitTenant.RouteClient(app)
	}
}
