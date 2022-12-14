package role

import (
	"github.com/bonkzero404/gaskn/config"
	"github.com/bonkzero404/gaskn/driver"
	"github.com/bonkzero404/gaskn/features/role/handlers"
	implements2 "github.com/bonkzero404/gaskn/features/role/interactors/implements"
	"github.com/bonkzero404/gaskn/features/role/repositories/implements"

	"github.com/gofiber/fiber/v2"
)

// RegisterFeature /*
func RegisterFeature(app *fiber.App) {

	roleRepository := implements.NewRoleRepository(driver.DB)
	roleService := implements2.NewRole(roleRepository)
	RoleHandler := handlers.NewRoleHandler(roleService)

	roleClientRepository := implements.NewRoleClientRepository(driver.DB)
	roleClientService := implements2.NewRoleClient(roleClientRepository, roleRepository)
	RoleClientHandler := handlers.NewRoleClientHandler(roleClientService)

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
		routesInitTenant.Route(app)
	}
}
