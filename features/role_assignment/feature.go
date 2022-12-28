package role_assignment

import (
	"github.com/bonkzero404/gaskn/config"
	"github.com/bonkzero404/gaskn/driver"
	roleRepo "github.com/bonkzero404/gaskn/features/role/repositories/implements"
	"github.com/bonkzero404/gaskn/features/role_assignment/handlers"
	roleAssignInteract "github.com/bonkzero404/gaskn/features/role_assignment/interactors/implements"
	roleAssignRepo "github.com/bonkzero404/gaskn/features/role_assignment/repositories/implements"
	"github.com/gofiber/fiber/v2"
)

// RegisterFeature /*
func RegisterFeature(app *fiber.App) {
	roleRepository := roleRepo.NewRoleRepository(driver.DB)
	roleClientRepository := roleRepo.NewRoleClientRepository(driver.DB)
	roleAssignmentRepository := roleAssignRepo.NewRoleAssignmentRepository(driver.DB)

	roleAssignment := roleAssignInteract.NewRoleAssignment(roleClientRepository, roleRepository, roleAssignmentRepository)
	RoleAssignmentHandler := handlers.NewRoleAssignmentHandler(roleAssignment)

	var routesInitTenant = ApiRoute{}

	routesInit := ApiRoute{
		RoleAssignmentHandler: *RoleAssignmentHandler,
	}

	routesInit.Route(app)

	/////////////////////////
	// If tenant is enabled
	/////////////////////////
	if config.Config("TENANCY") == "true" {
		routesInitTenant = ApiRoute{
			RoleAssignmentHandler: *RoleAssignmentHandler,
		}
		routesInitTenant.RouteClient(app)
	}

}
