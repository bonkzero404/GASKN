package role

import (
	"go-starterkit-project/app/middleware"
	"go-starterkit-project/modules/role/handlers"
	"go-starterkit-project/utils"

	"github.com/gofiber/fiber/v2"
)

type ApiRoute struct {
	RoleHandler handlers.RoleHandler
}

type ApiRouteClient struct {
	RoleClientHandler handlers.RoleClientHandler
}

// /////////////////
// Route Role
// /////////////////
func (handler *ApiRoute) Route(app fiber.Router) {
	const endpointGroup string = "/role"

	role := app.Group(utils.SetupApiGroup() + endpointGroup)
	feature := utils.RouteFeature{}

	role.Post(
		"/",
		middleware.Authenticate(),
		middleware.RateLimiter(5, 30),
		middleware.Permission(),
		handler.RoleHandler.CreateRole,
	).Name(
		feature.
			SetGroup("Role").
			SetName("CreateRole").
			SetDescription("Users can create roles").
			SetOnlyAdmin(true).
			Exec(),
	)

	role.Get(
		"/",
		middleware.Authenticate(),
		middleware.RateLimiter(5, 30),
		middleware.Permission(),
		handler.RoleHandler.GetRoleList,
	).Name(
		feature.
			SetGroup("Role").
			SetName("GetRoleLists").
			SetDescription("Users can get role lists").
			SetOnlyAdmin(true).
			Exec(),
	)

	role.Put(
		"/:id",
		middleware.Authenticate(),
		middleware.RateLimiter(5, 30),
		middleware.Permission(),
		handler.RoleHandler.UpdateRole,
	).Name(
		feature.
			SetGroup("Role").
			SetName("UpdateRole").
			SetDescription("Users can update roles").
			SetOnlyAdmin(true).
			Exec(),
	)

	role.Delete(
		"/:id",
		middleware.Authenticate(),
		middleware.RateLimiter(5, 30),
		middleware.Permission(),
		handler.RoleHandler.DeleteRole,
	).Name(
		feature.
			SetGroup("Role").
			SetName("UpdateRole").
			SetDescription("Users can delete roles").
			SetOnlyAdmin(true).
			Exec(),
	)

}

// /////////////////
// Route Role Client
// /////////////////
func (handler *ApiRouteClient) Route(app fiber.Router) {
	const endpointGroup string = "/role"

	roleClient := app.Group(utils.SetupSubApiGroup() + endpointGroup)
	feature := utils.RouteFeature{}

	roleClient.Post(
		"/",
		middleware.Authenticate(),
		middleware.RateLimiter(5, 30),
		middleware.Permission(),
		handler.RoleClientHandler.CreateClientRole,
	).Name(
		feature.
			SetGroup("RoleClient").
			SetName("CreateClientRole").
			SetDescription("Users (clients) can create roles").
			SetTenant(true).
			Exec(),
	)

	roleClient.Get(
		"/",
		middleware.Authenticate(),
		middleware.RateLimiter(5, 30),
		middleware.Permission(),
		handler.RoleClientHandler.GetRoleClientList,
	).Name(
		feature.
			SetGroup("RoleClient").
			SetName("GetClientRoleList").
			SetDescription("Users (clients) can get role lists").
			SetTenant(true).
			Exec(),
	)

	roleClient.Put(
		"/:id",
		middleware.Authenticate(),
		middleware.RateLimiter(5, 30),
		middleware.Permission(),
		handler.RoleClientHandler.UpdateRoleClient,
	).Name(
		feature.
			SetGroup("RoleClient").
			SetName("UpdateClientRole").
			SetDescription("Users (clients) can update role").
			SetTenant(true).
			Exec(),
	)

	roleClient.Delete(
		"/:id",
		middleware.Authenticate(),
		middleware.RateLimiter(5, 30),
		middleware.Permission(),
		handler.RoleClientHandler.DeleteRoleClient,
	).Name(
		feature.
			SetGroup("RoleClient").
			SetName("DeleteClientRole").
			SetDescription("Users (clients) can delete role").
			SetTenant(true).
			Exec(),
	)

}
