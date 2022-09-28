package user

import (
	"go-starterkit-project/app/middleware"
	"go-starterkit-project/modules/user/handlers"
	"go-starterkit-project/utils"

	"github.com/gofiber/fiber/v2"
)

type ApiRoute struct {
	UserHandler handlers.UserHandler
}

func (handler *ApiRoute) Route(app fiber.Router) {
	const endpointGroup string = "/user"

	user := app.Group(utils.SetupApiGroup() + endpointGroup)

	user.Post(
		"/register",
		middleware.RateLimiter(5, 30),
		handler.UserHandler.RegisterUser,
	).Name("RegisterUser")

	user.Post(
		"/activation",
		middleware.RateLimiter(5, 30),
		handler.UserHandler.UserActivation,
	).Name("UserActivation")

	user.Post(
		"/activation/re-send",
		middleware.RateLimiter(5, 30),
		handler.UserHandler.ReCreateUserActivation,
	).Name("ReCreateUserActivation")

	user.Post(
		"/request-forgot-password",
		middleware.RateLimiter(5, 30),
		handler.UserHandler.CreateActivationForgotPassword,
	).Name("CreateActivationForgotPassword")

	user.Post(
		"/forgot-password",
		middleware.RateLimiter(5, 30),
		handler.UserHandler.UpdatePassword,
	).Name("UpdatePassword")
}
