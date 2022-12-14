package auth

import (
	"github.com/bonkzero404/gaskn/driver"
	"github.com/bonkzero404/gaskn/features/auth/handlers"
	implements2 "github.com/bonkzero404/gaskn/features/auth/interactors/implements"
	roleRepo "github.com/bonkzero404/gaskn/features/role/repositories/implements"
	"github.com/bonkzero404/gaskn/features/user/repositories/implements"

	"github.com/gofiber/fiber/v2"
)

// RegisterFeature /*
func RegisterFeature(app *fiber.App) {
	userRepository := implements.NewUserRepository(driver.DB)
	roleRepository := roleRepo.NewRoleRepository(driver.DB)
	authService := implements2.NewAuth(userRepository, roleRepository)
	authHandler := handlers.NewAuthHandler(authService)

	routesInit := ApiRoute{
		AuthHandler: *authHandler,
	}

	routesInit.Route(app)
}
