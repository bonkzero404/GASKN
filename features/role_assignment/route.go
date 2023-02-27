package role_assignment

import (
	"github.com/bonkzero404/gaskn/app/http/builder"
	"github.com/bonkzero404/gaskn/app/http/middleware"
	"github.com/bonkzero404/gaskn/app/utils"
	"github.com/bonkzero404/gaskn/config"
	"github.com/gofiber/fiber/v2"

	"github.com/bonkzero404/gaskn/features/role_assignment/handlers"
)

type ApiRoute struct {
	RoleAssignmentHandler handlers.RoleAssignmentHandler
}

func (handler *ApiRoute) Route(app fiber.Router) {
	const endpointGroup string = "/role-assignment"

	role := builder.RouteBuilder{}

	role.Set(app).
		Group(utils.ApiBasePath() + endpointGroup).
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

	roleClient := builder.RouteBuilder{}

	roleClient.Set(app).
		Group(utils.ApiClientBasePath() + endpointGroup).
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
