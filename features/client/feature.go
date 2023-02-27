package role

import (
	"github.com/bonkzero404/gaskn/config"
	"github.com/bonkzero404/gaskn/features/client/handlers"
	clientInteract "github.com/bonkzero404/gaskn/features/client/interactors/implements"
	clientRepo "github.com/bonkzero404/gaskn/features/client/repositories/implements"
	userRepo "github.com/bonkzero404/gaskn/features/user/repositories/implements"
	"github.com/bonkzero404/gaskn/infrastructures"

	"github.com/gofiber/fiber/v2"
)

// RegisterFeature /*
func RegisterFeature(app *fiber.App) {
	userReposiry := userRepo.NewUserRepository(infrastructures.DB)
	clientRepository := clientRepo.NewClientRepository(infrastructures.DB)
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
