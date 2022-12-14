package role_assignment

import (
	"github.com/bonkzero404/gaskn/driver"
	"github.com/bonkzero404/gaskn/features/role/repositories/implements"
	"github.com/bonkzero404/gaskn/features/role_assignment/handlers"
	implements2 "github.com/bonkzero404/gaskn/features/role_assignment/interactors/implements"

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
