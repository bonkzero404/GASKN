package user

import (
	"github.com/gofiber/fiber/v2"

	"gaskn/app/middleware"
	"gaskn/features/user/handlers"
	"gaskn/utils"
)

type ApiRoute struct {
	UserHandler handlers.UserHandler
}

type ApiRouteClient struct {
	UserClientHandler handlers.UserClientHandler
}

func (handler *ApiRoute) Route(app fiber.Router) {
	const endpointGroup string = "/user"

	user := app.Group(utils.SetupApiGroup() + endpointGroup)

	user.Post(
		"/register",
		middleware.RateLimiter(5, 30),
		handler.UserHandler.RegisterUser,
	)

	user.Post(
		"/activation",
		middleware.RateLimiter(5, 30),
		handler.UserHandler.UserActivation,
	)

	user.Post(
		"/activation/re-send",
		middleware.RateLimiter(5, 30),
		handler.UserHandler.ReCreateUserActivation,
	)

	user.Post(
		"/request-forgot-password",
		middleware.RateLimiter(5, 30),
		handler.UserHandler.CreateActivationForgotPassword,
	)

	user.Post(
		"/forgot-password",
		middleware.RateLimiter(5, 30),
		handler.UserHandler.UpdatePassword,
	)
}

// /////////////////
// Route Role Client
// /////////////////
func (handler *ApiRouteClient) Route(app fiber.Router) {
	const endpointGroup string = "/user"

	userClient := app.Group(utils.SetupSubApiGroup() + endpointGroup)
	feature := utils.RouteFeature{}

	userClient.Post(
		"/invitation",
		middleware.Authenticate(),
		middleware.RateLimiter(5, 30),
		middleware.Permission(),
		handler.UserClientHandler.CreateUserInvitation,
	).Name(
		feature.
			SetGroup("Client/UserInvitation").
			SetName("CreateClientUserInvitation").
			SetDescription("User can invite other users to join organizations").
			SetTenant(true).
			Exec(),
	)

	userClient.Post(
		"/invitation/acceptance",
		middleware.Authenticate(),
		middleware.RateLimiter(5, 30),
		middleware.Permission(),
		handler.UserClientHandler.UserInvitationAcceptance,
	).Name(
		feature.
			SetGroup("Client/UserInvitationAcceptance").
			SetName("AcceptanceClientUserInvitation").
			SetDescription("User can accept invitation to join organizations").
			SetTenant(true).
			Exec(),
	)
}
