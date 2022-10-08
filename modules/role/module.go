package role

import (
	"go-starterkit-project/config"
	"go-starterkit-project/database/driver"
	"go-starterkit-project/modules/role/handlers"
	"go-starterkit-project/modules/role/repositories"
	"go-starterkit-project/modules/role/services"

	"github.com/gofiber/fiber/v2"
)

/*
*
This function is for registering repository - service - handler
*/
func RegisterModule(app *fiber.App) {

	roleRepository := repositories.NewRoleRepository(driver.DB)
	roleService := services.NewRoleService(roleRepository)
	RoleHandler := handlers.NewRoleHandler(roleService)

	roleClientRepository := repositories.NewRoleClientRepository(driver.DB)
	roleClientService := services.NewRoleClientService(roleClientRepository, roleRepository)
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
