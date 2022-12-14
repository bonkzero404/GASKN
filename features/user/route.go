package user

import (
	"github.com/gofiber/fiber/v2"

	"github.com/bonkzero404/gaskn/app/middleware"
	"github.com/bonkzero404/gaskn/features/user/handlers"
	"github.com/bonkzero404/gaskn/utils"
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
		SetRouteDescriptionKeyLang("route:client:user:invitation").
		SetRouteTenant(true).
		Execute()

	userClient.Post(
		"/invitation/acceptance",
		middleware.Authenticate(),
		middleware.RateLimiter(5, 30),
		handler.UserClientHandler.UserInvitationAcceptance,
	)
}
