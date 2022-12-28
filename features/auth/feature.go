package auth

import (
	"github.com/bonkzero404/gaskn/driver"
	"github.com/bonkzero404/gaskn/features/auth/handlers"
	"github.com/bonkzero404/gaskn/features/auth/interactors/implements"
	roleRepo "github.com/bonkzero404/gaskn/features/role/repositories/implements"
	userRepo "github.com/bonkzero404/gaskn/features/user/repositories/implements"

	"github.com/gofiber/fiber/v2"
)

// RegisterFeature /*
func RegisterFeature(app *fiber.App) {
	userRepository := userRepo.NewUserRepository(driver.DB)
	roleRepository := roleRepo.NewRoleRepository(driver.DB)
	auth := implements.NewAuth(userRepository, roleRepository)
	authHandler := handlers.NewAuthHandler(auth)

	routesInit := ApiRoute{
		AuthHandler: *authHandler,
	}

	routesInit.Route(app)
}
