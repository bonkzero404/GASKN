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

	roleClient := app.Group(utils.SetupSubApiGroup() + endpointGroup)
	feature := utils.RouteFeature{}

	roleClient.Post(
		"/",
		middleware.Authenticate(),
		middleware.RateLimiter(5, 30),
		middleware.Permission(),
		handler.RoleAssignmentHandler.CreateRoleAssignment,
	).Name(
		feature.
			SetGroup("Client/Role/Assignment").
			SetName("CreateClientRoleAssignment").
			SetDescription("Users (clients) can assignment roles").
			SetTenant(true).
			Exec(),
	)

	roleClient.Delete(
		"/",
		middleware.Authenticate(),
		middleware.RateLimiter(5, 30),
		middleware.Permission(),
		handler.RoleAssignmentHandler.RemoveRoleAssignment,
	).Name(
		feature.
			SetGroup("Client/Role/Assignment").
			SetName("RemoveClientRoleAssignment").
			SetDescription("Users (clients) can remove assignment roles").
			SetTenant(true).
			Exec(),
	)
}
