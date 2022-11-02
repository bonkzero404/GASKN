package role

import (
	"github.com/gofiber/fiber/v2"

	"gaskn/app/middleware"
	"gaskn/modules/client/handlers"
	"gaskn/utils"
)

type ApiRoute struct {
	ClientHandler handlers.ClientHandler
}

func (handler *ApiRoute) Route(app fiber.Router) {
	const endpointGroup string = "/client"

	client := app.Group(utils.SetupApiGroup() + endpointGroup)
	clientAcc := app.Group(utils.SetupSubApiGroup())
	feature := utils.RouteFeature{}

	client.Post(
		"/",
		middleware.Authenticate(),
		middleware.RateLimiter(5, 30),
		handler.ClientHandler.CreateClient,
	).Name(
		feature.
			SetGroup("Client").
			SetName("CreateClient").
			SetDescription("Users can create client").
			Exec(),
	)

	client.Get(
		"/user",
		middleware.Authenticate(),
		middleware.RateLimiter(5, 30),
		handler.ClientHandler.GetClientByUser,
	).Name(
		feature.
			SetGroup("Client").
			SetName("GetClientByUser").
			SetDescription("Users can get client by user").
			Exec(),
	)

	clientAcc.Put(
		"/update",
		middleware.Authenticate(),
		middleware.RateLimiter(5, 30),
		middleware.Permission(),
		handler.ClientHandler.UpdateClient,
	).Name(feature.
		SetGroup("Client").
		SetName("UpdateClient").
		SetDescription("Users can update client").
		Exec())

}
