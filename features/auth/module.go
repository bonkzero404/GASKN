package auth

import (
	"gaskn/database/driver"
	"gaskn/features/auth/handlers"
	"gaskn/features/auth/services"
	roleRepo "gaskn/features/role/repositories"
	"gaskn/features/user/repositories"

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
