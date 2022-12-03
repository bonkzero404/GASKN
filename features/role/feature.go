package role

import (
	"gaskn/config"
	"gaskn/driver"
	"gaskn/features/role/handlers"
	"gaskn/features/role/repositories"
	"gaskn/features/role/services"

	"github.com/gofiber/fiber/v2"
)

// RegisterFeature /*
func RegisterFeature(app *fiber.App) {

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
