package role

import (
	"github.com/bonkzero404/gaskn/config"
	"github.com/gofiber/fiber/v2"

	"github.com/bonkzero404/gaskn/app/middleware"
	"github.com/bonkzero404/gaskn/features/client/handlers"
	"github.com/bonkzero404/gaskn/utils"
)

type ApiRoute struct {
	ClientHandler handlers.ClientHandler
}

type ApiRouteClient struct {
	ClientHandler handlers.ClientHandler
}

func (handler *ApiRoute) Route(app fiber.Router) {
	var endpointGroup = "/" + config.Config("API_CLIENT")

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
		SetRouteDescriptionKeyLang("route:client:update").
		SetRouteTenant(true).
		Execute()

}
