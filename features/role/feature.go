package role

import (
	"github.com/bonkzero404/gaskn/config"
	"github.com/bonkzero404/gaskn/driver"
	"github.com/bonkzero404/gaskn/features/role/handlers"
	roleInteract "github.com/bonkzero404/gaskn/features/role/interactors/implements"
	roleRepo "github.com/bonkzero404/gaskn/features/role/repositories/implements"

	"github.com/gofiber/fiber/v2"
)

// RegisterFeature /*
func RegisterFeature(app *fiber.App) {

	roleRepository := roleRepo.NewRoleRepository(driver.DB)
	role := roleInteract.NewRole(roleRepository)
	RoleHandler := handlers.NewRoleHandler(role)

	roleClientRepository := roleRepo.NewRoleClientRepository(driver.DB)
	roleClient := roleInteract.NewRoleClient(roleClientRepository, roleRepository)
	RoleClientHandler := handlers.NewRoleClientHandler(roleClient)

	var routesInitTenant = ApiRouteClient{}

	routesInit := ApiRoute{
		RoleHandler: *RoleHandler,
	}

	routesInit.Route(app)

	/////////////////////////
	// If tenant is enabled
	/////////////////////////
	if config.Config("TENANCY") == "true" {
		routesInitTenant = ApiRouteClient{
			RoleClientHandler: *RoleClientHandler,
		}
		routesInitTenant.RouteClient(app)
	}
}
