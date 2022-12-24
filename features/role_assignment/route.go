package role_assignment

import (
	"github.com/bonkzero404/gaskn/config"
	"github.com/gofiber/fiber/v2"

	"github.com/bonkzero404/gaskn/app/middleware"
	"github.com/bonkzero404/gaskn/features/role_assignment/handlers"
	"github.com/bonkzero404/gaskn/utils"
)

type ApiRoute struct {
	RoleAssignmentHandler handlers.RoleAssignmentHandler
}

func (handler *ApiRoute) Route(app fiber.Router) {
	const endpointGroup string = "/role-assignment"

	role := utils.GasknRouter{}

	role.Set(app).
		Group(utils.SetupApiGroup() + endpointGroup).
		SetGroupName("Role/Assignment")

	role.Post(
		"/",
		middleware.Authenticate(),
		middleware.RateLimiter(5, 30),
		middleware.Permission(),
		handler.RoleAssignmentHandler.CreateRoleAssignment,
	).
		SetRouteName("CreateRoleAssignment").
		SetRouteDescriptionKeyLang(config.RouteClientRoleAssignmentAdd).
		Execute()

	role.Delete(
		"/",
		middleware.Authenticate(),
		middleware.RateLimiter(5, 30),
		middleware.Permission(),
		handler.RoleAssignmentHandler.RemoveRoleAssignment,
	).
		SetRouteName("RemoveRoleAssignment").
		SetRouteDescriptionKeyLang(config.RouteCLientRoleAssignmentDelete).
		Execute()

	role.Post(
		"/user",
		middleware.Authenticate(),
		middleware.RateLimiter(5, 30),
		middleware.Permission(),
		handler.RoleAssignmentHandler.AssignUserPermission,
	).
		SetRouteName("CreateUserRoleAssignment").
		SetRouteDescriptionKeyLang(config.RouteClientRoleAssignment).
		Execute()

	role.Get(
		"/list/:RoleId",
		middleware.Authenticate(),
		middleware.RateLimiter(5, 30),
		middleware.Permission(),
		handler.RoleAssignmentHandler.GetPermissionByRole,
	).
		SetRouteName("GetPermissionRoleAssignment").
		SetRouteDescriptionKeyLang(config.RouteClientRoleViewAssignment).
		Execute()
}

func (handler *ApiRoute) RouteClient(app fiber.Router) {
	const endpointGroup string = "/role-assignment"

	roleClient := utils.GasknRouter{}

	roleClient.Set(app).
		Group(utils.SetupSubApiGroup() + endpointGroup).
		SetGroupName("Client/Role/Assignment")

	roleClient.Post(
		"/",
		middleware.Authenticate(),
		middleware.RateLimiter(5, 30),
		middleware.Permission(),
		handler.RoleAssignmentHandler.CreateRoleAssignment,
	).
		SetRouteName("CreateClientRoleAssignment").
		SetRouteDescriptionKeyLang(config.RouteClientRoleAssignmentAdd).
		SetRouteTenant(true).
		Execute()

	roleClient.Delete(
		"/",
		middleware.Authenticate(),
		middleware.RateLimiter(5, 30),
		middleware.Permission(),
		handler.RoleAssignmentHandler.RemoveRoleAssignment,
	).
		SetRouteName("RemoveClientRoleAssignment").
		SetRouteDescriptionKeyLang(config.RouteCLientRoleAssignmentDelete).
		SetRouteTenant(true).
		Execute()

	roleClient.Post(
		"/user",
		middleware.Authenticate(),
		middleware.RateLimiter(5, 30),
		middleware.Permission(),
		handler.RoleAssignmentHandler.AssignUserPermission,
	).
		SetRouteName("CreateUserClientRoleAssignment").
		SetRouteDescriptionKeyLang(config.RouteClientRoleAssignment).
		SetRouteTenant(true).
		Execute()

	roleClient.Get(
		"/list/:RoleId",
		middleware.Authenticate(),
		middleware.RateLimiter(5, 30),
		middleware.Permission(),
		handler.RoleAssignmentHandler.GetPermissionByRole,
	).
		SetRouteName("GetPermissionClientRoleAssignment").
		SetRouteDescriptionKeyLang(config.RouteClientRoleViewAssignment).
		SetRouteTenant(true).
		Execute()
}
