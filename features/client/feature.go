package role

import (
	"github.com/bonkzero404/gaskn/config"
	"github.com/bonkzero404/gaskn/driver"
	"github.com/bonkzero404/gaskn/features/client/handlers"
	"github.com/bonkzero404/gaskn/features/client/interactors/implements"
	implements2 "github.com/bonkzero404/gaskn/features/client/repositories/implements"
	userRepo "github.com/bonkzero404/gaskn/features/user/repositories/implements"

	"github.com/gofiber/fiber/v2"
)

// RegisterFeature /*
func RegisterFeature(app *fiber.App) {
	userReposiry := userRepo.NewUserRepository(driver.DB)
	clientRepository := implements2.NewClientRepository(driver.DB)
	clientService := implements.NewClient(clientRepository, userReposiry)
	ClientHandler := handlers.NewClientHandler(clientService)

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
		routesInitTenant.Route(app)
	}
}
