package role

import (
	"github.com/bonkzero404/gaskn/app/http/builder"
	"github.com/bonkzero404/gaskn/app/http/middleware"
	"github.com/bonkzero404/gaskn/app/utils"
	"github.com/bonkzero404/gaskn/config"
	"github.com/gofiber/fiber/v2"

	"github.com/bonkzero404/gaskn/features/role/handlers"
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

	var role = builder.RouteBuilder{}

	role.Set(app).
		Group(utils.ApiBasePath() + endpointGroup).
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

// RouteClient /**
func (handler *ApiRouteClient) RouteClient(app fiber.Router) {
	const endpointGroup string = "/role"

	var roleClient = builder.RouteBuilder{}

	roleClient.Set(app).
		Group(utils.ApiClientBasePath() + endpointGroup).
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
