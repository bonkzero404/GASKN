package auth

import (
	"go-starterkit-project/database/driver"
	"go-starterkit-project/modules/auth/handlers"
	"go-starterkit-project/modules/auth/services"
	roleRepo "go-starterkit-project/modules/role/repositories"
	"go-starterkit-project/modules/user/repositories"

	"github.com/gofiber/fiber/v2"
)

/*
*
This function is for registering repository - service - handler
*/
func RegisterModule(app *fiber.App) {
	userRepository := repositories.NewUserRepository(driver.DB)
	roleRepository := roleRepo.NewRoleRepository(driver.DB)
	authService := services.NewAuthService(userRepository, roleRepository)
	authHandler := handlers.NewAuthHandler(authService)

	routesInit := ApiRoute{
		AuthHandler: *authHandler,
	}

	routesInit.Route(app)
}
