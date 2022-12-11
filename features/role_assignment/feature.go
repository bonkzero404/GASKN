package role_assignment

import (
	"gaskn/driver"
	"gaskn/features/role/repositories/implements"
	"gaskn/features/role_assignment/handlers"
	implements2 "gaskn/features/role_assignment/interactors/implements"

	"github.com/gofiber/fiber/v2"
)

// RegisterFeature /*
func RegisterFeature(app *fiber.App) {
	roleRepository := implements.NewRoleRepository(driver.DB)
	roleClientRepository := implements.NewRoleClientRepository(driver.DB)

	roleAssignmentService := implements2.NewRoleAssignment(roleClientRepository, roleRepository)
	RoleAssignmentHandler := handlers.NewRoleAssignmentHandler(roleAssignmentService)

	routesInit := ApiRouteClient{
		RoleAssignmentHandler: *RoleAssignmentHandler,
	}

	routesInit.Route(app)

}
