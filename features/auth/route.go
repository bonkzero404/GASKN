package auth

import (
	"github.com/bonkzero404/gaskn/app/middleware"
	"github.com/bonkzero404/gaskn/features/auth/handlers"
	"github.com/bonkzero404/gaskn/utils"

	"github.com/gofiber/fiber/v2"
)

type ApiRoute struct {
	AuthHandler handlers.AuthHandler
}

func (handler *ApiRoute) Route(app fiber.Router) {
	const endpointGroup string = "/auth"

	user := utils.GasknRouter{}

	user.Set(app).Group(utils.SetupApiGroup() + endpointGroup)

	user.Post(
		"/",
		middleware.RateLimiter(5, 120),
		handler.AuthHandler.Authentication,
	)

	user.Get(
		"/me",
		middleware.Authenticate(),
		middleware.RateLimiter(200, 30),
		handler.AuthHandler.GetProfile,
	)

	user.Get(
		"/refresh-token",
		middleware.Authenticate(),
		middleware.RateLimiter(200, 30),
		handler.AuthHandler.RefreshToken,
	)
}
