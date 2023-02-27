package role

import (
	"github.com/bonkzero404/gaskn/app/http/builder"
	"github.com/bonkzero404/gaskn/app/http/middleware"
	"github.com/bonkzero404/gaskn/app/utils"
	"github.com/bonkzero404/gaskn/config"
	"github.com/gofiber/fiber/v2"

	"github.com/bonkzero404/gaskn/features/client/handlers"
)

type ApiRoute struct {
	ClientHandler handlers.ClientHandler
}

type ApiRouteClient struct {
	ClientHandler handlers.ClientHandler
}

func (handler *ApiRoute) Route(app fiber.Router) {
	var endpointGroup = "/" + config.Config("API_CLIENT")

	client := builder.RouteBuilder{}
	client.Set(app).Group(utils.ApiBasePath() + endpointGroup)

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

func (handler *ApiRouteClient) RouteClient(app fiber.Router) {
	clientAcc := builder.RouteBuilder{}
	clientAcc.Set(app).Group(utils.ApiClientBasePath()).SetGroupName("Client") //app.Group(utils.ApiClientBasePath())

	clientAcc.Put(
		"/update",
		middleware.Authenticate(),
		middleware.RateLimiter(5, 30),
		middleware.Permission(),
		handler.ClientHandler.UpdateClient,
	).
		SetRouteName("UpdateClient").
		SetRouteDescriptionKeyLang(config.RouteClientUpdate).
		SetRouteTenant(true).
		Execute()

}
