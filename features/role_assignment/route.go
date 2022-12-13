package role_assignment

import (
	"github.com/gofiber/fiber/v2"

	"gaskn/app/middleware"
	"gaskn/features/role_assignment/handlers"
	"gaskn/utils"
)

type ApiRouteClient struct {
	RoleAssignmentHandler handlers.RoleAssignmentHandler
}

func (handler *ApiRouteClient) Route(app fiber.Router) {
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
