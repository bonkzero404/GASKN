package user

import (
	"github.com/bonkzero404/gaskn/app/http/builder"
	"github.com/bonkzero404/gaskn/app/http/middleware"
	"github.com/bonkzero404/gaskn/app/utils"
	"github.com/bonkzero404/gaskn/config"
	"github.com/gofiber/fiber/v2"

	"github.com/bonkzero404/gaskn/features/user/handlers"
)

type ApiRoute struct {
	UserHandler       handlers.UserHandler
	UserClientHandler handlers.UserClientHandler
}

func (handler *ApiRoute) Route(app fiber.Router) {
	const endpointGroup string = "/user"

	user := builder.RouteBuilder{}
	user.Set(app).Group(utils.ApiBasePath() + endpointGroup).SetGroupName("User")

	user.Post(
		"/register",
		middleware.RateLimiter(5, 30),
		handler.UserHandler.CreateUser,
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
		"/:CreateUser",
		middleware.Authenticate(),
		middleware.RateLimiter(5, 30),
		middleware.Permission(),
		handler.UserHandler.CreateUser,
	).
		SetRouteName("CreateUser").
		SetRouteDescriptionKeyLang("user:create").
		Execute()

}

func (handler *ApiRoute) RouteClient(app fiber.Router) {
	const endpointGroup string = "/user"

	userClient := builder.RouteBuilder{}
	userClient.Set(app).
		Group(utils.ApiClientBasePath() + endpointGroup).
		SetGroupName("Client/UserInvitation") // app.Group(utils.ApiClientBasePath() + endpointGroup)

	userClient.Post(
		"/:CreateUser",
		middleware.Authenticate(),
		middleware.RateLimiter(5, 30),
		middleware.Permission(),
		handler.UserHandler.CreateUser,
	).
		SetRouteName("CreateClientUser").
		SetRouteDescriptionKeyLang(config.RouteUserCreate).
		SetRouteTenant(true).
		Execute()

	userClient.Post(
		"/invitation",
		middleware.Authenticate(),
		middleware.RateLimiter(5, 30),
		middleware.Permission(),
		handler.UserClientHandler.CreateUserInvitation,
	).
		SetRouteName("CreateClientUserInvitation").
		SetRouteDescriptionKeyLang(config.RouteClientUserInvitation).
		SetRouteTenant(true).
		Execute()

	userClient.Post(
		"/invitation/acceptance",
		middleware.Authenticate(),
		middleware.RateLimiter(5, 30),
		handler.UserClientHandler.UserInvitationAcceptance,
	)
}
