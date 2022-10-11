package auth

import (
	"gaskn/database/driver"
	"gaskn/modules/auth/handlers"
	"gaskn/modules/auth/services"
	roleRepo "gaskn/modules/role/repositories"
	"gaskn/modules/user/repositories"

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
