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

func (handler *ApiRoute) Route(app fiber.Router) {
	const endpointGroup string = "/role"

	role := app.Group(utils.SetupApiGroup() + endpointGroup)

	role.Post("/", middleware.Authenticate(), middleware.RateLimiter(5, 30), handler.RoleHandler.CreateRole)

	role.Get("/", middleware.Authenticate(), middleware.RateLimiter(5, 30), handler.RoleHandler.GetRoleList)

	role.Put("/:id", middleware.Authenticate(), middleware.RateLimiter(5, 30), handler.RoleHandler.UpdateRole)

	role.Delete("/:id", middleware.Authenticate(), middleware.RateLimiter(5, 30), handler.RoleHandler.DeleteRole)

}
