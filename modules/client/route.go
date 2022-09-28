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
	roleClient := app.Group(utils.SetupSubApiGroup())

	role.Post(
		"/",
		middleware.Authenticate(),
		middleware.RateLimiter(5, 30),
		handler.ClientHandler.CreateClient,
	).Name("CreateClient")

	role.Get(
		"/user",
		middleware.Authenticate(),
		middleware.RateLimiter(5, 30),
		handler.ClientHandler.GetClientByUser,
	).Name("GetClientByUser")

	roleClient.Put(
		"/update",
		middleware.Authenticate(),
		middleware.RateLimiter(5, 30),
		middleware.Permission(),
		handler.ClientHandler.UpdateClient,
	).Name("UpdateClient")

}
