package role

import (
	"github.com/gofiber/fiber/v2"

	"gaskn/app/middleware"
	"gaskn/features/role/handlers"
	"gaskn/utils"
)

type ApiRoute struct {
	RoleHandler handlers.RoleHandler
}

type ApiRouteClient struct {
	RoleClientHandler handlers.RoleClientHandler
}

// Route /**
func (handler *ApiRoute) Route(app fiber.Router) {
	const endpointGroup string = "/role"

	var role = utils.GasknRouter{}

	role.Set(app).
		Group(utils.SetupApiGroup() + endpointGroup).
		SetGroupName("Role")

	role.Post(
		"/",
		middleware.Authenticate(),
		middleware.RateLimiter(5, 30),
		middleware.Permission(),
		handler.RoleHandler.CreateRole,
	).
		SetRouteName("CreateRole").
		SetRouteDescription("User can create roles").
		Execute()

	role.Get(
		"/",
		middleware.Authenticate(),
		middleware.RateLimiter(5, 30),
		middleware.Permission(),
		handler.RoleHandler.GetRoleList,
	).
		SetRouteName("GetRoleLists").
		SetRouteDescription("Users can get role lists").
		Execute()

	role.Put(
		"/:id",
		middleware.Authenticate(),
		middleware.RateLimiter(5, 30),
		middleware.Permission(),
		handler.RoleHandler.UpdateRole,
	).
		SetRouteName("UpdateRole").
		SetRouteDescription("Users can update roles").
		Execute()

	role.Delete(
		"/:id",
		middleware.Authenticate(),
		middleware.RateLimiter(5, 30),
		middleware.Permission(),
		handler.RoleHandler.DeleteRole,
	).
		SetRouteName("DeleteRole").
		SetRouteDescription("Users can delete roles").
		Execute()

}

// Route /**
func (handler *ApiRouteClient) Route(app fiber.Router) {
	const endpointGroup string = "/role"

	var roleClient = utils.GasknRouter{}

	roleClient.Set(app).
		Group(utils.SetupSubApiGroup() + endpointGroup).
		SetGroupName("Client/Role")

	roleClient.Post(
		"/",
		middleware.Authenticate(),
		middleware.RateLimiter(5, 30),
		middleware.Permission(),
		handler.RoleClientHandler.CreateClientRole,
	).
		SetRouteName("CreateClientRole").
		SetRouteDescription("Users (clients) can create roles").
		SetRouteTenant(true).
		Execute()

	roleClient.Get(
		"/",
		middleware.Authenticate(),
		middleware.RateLimiter(5, 30),
		middleware.Permission(),
		handler.RoleClientHandler.GetRoleClientList,
	).
		SetRouteName("GetClientRoleList").
		SetRouteDescription("Users (clients) can get role lists").
		SetRouteTenant(true).
		Execute()

	roleClient.Put(
		"/:id",
		middleware.Authenticate(),
		middleware.RateLimiter(5, 30),
		middleware.Permission(),
		handler.RoleClientHandler.UpdateRoleClient,
	).
		SetRouteName("UpdateClientRole").
		SetRouteDescription("Users (clients) can update role").
		SetRouteTenant(true).
		Execute()

	roleClient.Delete(
		"/:id",
		middleware.Authenticate(),
		middleware.RateLimiter(5, 30),
		middleware.Permission(),
		handler.RoleClientHandler.DeleteRoleClient,
	).
		SetRouteName("DeleteClientRole").
		SetRouteDescription("Users (clients) can delete role").
		SetRouteTenant(true).
		Execute()

}
