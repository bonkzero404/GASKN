package role_assignment

import (
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
		SetRouteDescriptionKeyLang("route:client:role:assignment:add").
		Execute()

	role.Delete(
		"/",
		middleware.Authenticate(),
		middleware.RateLimiter(5, 30),
		middleware.Permission(),
		handler.RoleAssignmentHandler.RemoveRoleAssignment,
	).
		SetRouteName("RemoveRoleAssignment").
		SetRouteDescriptionKeyLang("route:client:role:assignment:remove").
		Execute()

	role.Post(
		"/user",
		middleware.Authenticate(),
		middleware.RateLimiter(5, 30),
		middleware.Permission(),
		handler.RoleAssignmentHandler.AssignUserPermitToRole,
	).
		SetRouteName("CreateUserRoleAssignment").
		SetRouteDescriptionKeyLang("route:client:role:assignment:assign").
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
		SetRouteDescriptionKeyLang("route:client:role:assignment:add").
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
		SetRouteDescriptionKeyLang("route:client:role:assignment:remove").
		SetRouteTenant(true).
		Execute()

	roleClient.Post(
		"/user",
		middleware.Authenticate(),
		middleware.RateLimiter(5, 30),
		middleware.Permission(),
		handler.RoleAssignmentHandler.AssignUserPermitToRole,
	).
		SetRouteName("CreateUserClientRoleAssignment").
		SetRouteDescriptionKeyLang("route:client:role:assignment:assign").
		SetRouteTenant(true).
		Execute()
}
