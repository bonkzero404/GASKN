package role

import (
	"go-starterkit-project/app/middleware"
	"go-starterkit-project/modules/client/handlers"
	"go-starterkit-project/utils"

	"github.com/gofiber/fiber/v2"
)

type ApiRoute struct {
	ClientHandler handlers.ClientHandler
}

func (handler *ApiRoute) Route(app fiber.Router) {
	const endpointGroup string = "/client"

	role := app.Group(utils.SetupApiGroup() + endpointGroup)

	role.Post("/", middleware.Authenticate(), middleware.RateLimiter(5, 30), handler.ClientHandler.CreateClient)

	role.Put("/:id", middleware.Authenticate(), middleware.RateLimiter(5, 30), handler.ClientHandler.UpdateClient)

}
