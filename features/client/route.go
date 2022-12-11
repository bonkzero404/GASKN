package role

import (
	"github.com/gofiber/fiber/v2"

	"gaskn/app/middleware"
	"gaskn/features/client/handlers"
	"gaskn/utils"
)

type ApiRoute struct {
	ClientHandler handlers.ClientHandler
}

type ApiRouteClient struct {
	ClientHandler handlers.ClientHandler
}

func (handler *ApiRoute) Route(app fiber.Router) {
	const endpointGroup string = "/client"

	client := utils.GasknRouter{}
	client.Set(app).Group(utils.SetupApiGroup() + endpointGroup)

	client.Post(
		"/",
		middleware.Authenticate(),
		middleware.RateLimiter(5, 30),
		handler.ClientHandler.CreateClient,
	)

	client.Get(
		"/user",
		middleware.Authenticate(),
		middleware.RateLimiter(5, 30),
		handler.ClientHandler.GetClientByUser,
	)

}

func (handler *ApiRouteClient) Route(app fiber.Router) {
	clientAcc := utils.GasknRouter{}
	clientAcc.Set(app).Group(utils.SetupSubApiGroup()).SetGroupName("Client") //app.Group(utils.SetupSubApiGroup())

	clientAcc.Put(
		"/update",
		middleware.Authenticate(),
		middleware.RateLimiter(5, 30),
		middleware.Permission(),
		handler.ClientHandler.UpdateClient,
	).
		SetRouteName("UpdateClient").
		SetRouteDescription("Users can update client").
		SetRouteTenant(true).
		Execute()

}
