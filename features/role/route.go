package role

import (
	"github.com/bonkzero404/gaskn/config"
	"github.com/gofiber/fiber/v2"

	"github.com/bonkzero404/gaskn/app/middleware"
	"github.com/bonkzero404/gaskn/features/role/handlers"
	"github.com/bonkzero404/gaskn/utils"
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
		SetRouteDescriptionKeyLang(config.RouteRoleAdd).
		Execute()

	role.Get(
		"/",
		middleware.Authenticate(),
		middleware.RateLimiter(5, 30),
		middleware.Permission(),
		handler.RoleHandler.GetRoleList,
	).
		SetRouteName("GetRoleLists").
		SetRouteDescriptionKeyLang(config.RouteRoleList).
		Execute()

	role.Put(
		"/:id",
		middleware.Authenticate(),
		middleware.RateLimiter(5, 30),
		middleware.Permission(),
		handler.RoleHandler.UpdateRole,
	).
		SetRouteName("UpdateRole").
		SetRouteDescriptionKeyLang(config.RouteRoleUpdate).
		Execute()

	role.Delete(
		"/:id",
		middleware.Authenticate(),
		middleware.RateLimiter(5, 30),
		middleware.Permission(),
		handler.RoleHandler.DeleteRole,
	).
		SetRouteName("DeleteRole").
		SetRouteDescriptionKeyLang(config.RouteRoleDelete).
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
		SetRouteDescriptionKeyLang(config.RouteClientRoleAdd).
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
		SetRouteDescriptionKeyLang(config.RouteClientRoleList).
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
		SetRouteDescriptionKeyLang(config.RouteClientRoleUpdate).
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
		SetRouteDescriptionKeyLang(config.RouteClientRoleDelete).
		SetRouteTenant(true).
		Execute()

}
