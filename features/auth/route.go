package auth

import (
	"gaskn/app/middleware"
	"gaskn/features/auth/handlers"
	"gaskn/utils"

	"github.com/gofiber/fiber/v2"
)

type ApiRoute struct {
	AuthHandler handlers.AuthHandler
}

func (handler *ApiRoute) Route(app fiber.Router) {
	const endpointGroup string = "/auth"

	user := app.Group(utils.SetupApiGroup() + endpointGroup)

	user.Post(
		"/",
		middleware.RateLimiter(5, 120),
		handler.AuthHandler.Authentication,
	)

	user.Get(
		"/me",
		middleware.Authenticate(),
		middleware.RateLimiter(5, 30),
		handler.AuthHandler.GetProfile,
	)

	user.Get(
		"/refresh-token",
		middleware.Authenticate(),
		middleware.RateLimiter(5, 30),
		handler.AuthHandler.RefreshToken,
	)
}
