package user

import (
	"github.com/gofiber/fiber/v2"

	"gaskn/app/middleware"
	"gaskn/features/user/handlers"
	"gaskn/utils"
)

type ApiRoute struct {
	UserHandler       handlers.UserHandler
	UserClientHandler handlers.UserClientHandler
}

type ApiRouteClient struct {
	UserClientHandler handlers.UserClientHandler
}

func (handler *ApiRoute) Route(app fiber.Router) {
	const endpointGroup string = "/user"

	user := utils.GasknRouter{}
	user.Set(app).Group(utils.SetupApiGroup() + endpointGroup)

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

	user.Post(
		"/invitation/acceptance",
		middleware.Authenticate(),
		middleware.RateLimiter(5, 30),
		handler.UserClientHandler.UserInvitationAcceptance,
	)
}

func (handler *ApiRouteClient) Route(app fiber.Router) {
	const endpointGroup string = "/user"

	userClient := utils.GasknRouter{}
	userClient.Set(app).
		Group(utils.SetupSubApiGroup() + endpointGroup).
		SetGroupName("Client/UserInvitation") // app.Group(utils.SetupSubApiGroup() + endpointGroup)

	userClient.Post(
		"/invitation",
		middleware.Authenticate(),
		middleware.RateLimiter(5, 30),
		middleware.Permission(),
		handler.UserClientHandler.CreateUserInvitation,
	).
		SetRouteName("CreateClientUserInvitation").
		SetRouteDescription("User can invite other users to join organizations").
		SetRouteTenant(true).
		Execute()
}
