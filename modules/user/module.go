package user

import (
	"github.com/gofiber/fiber/v2"

	"gaskn/database/driver"
	"gaskn/modules/user/contracts"
	"gaskn/modules/user/handlers"
	"gaskn/modules/user/repositories"
	"gaskn/modules/user/services"
	"gaskn/modules/user/services/factories"
)

/*
*
Service factory registration
*/
func registerActivationFactory(userActivationRepository contracts.UserActionCodeRepository) factories.ActionFactoryInterface {
	actFactory := factories.NewUserActivationServiceFactory(userActivationRepository)
	forgotPassFactory := factories.NewUserForgotPassServiceFactory(userActivationRepository)

	return factories.NewActionFactory(actFactory, forgotPassFactory)
}

/*
*
This function is for registering repository - service - handler
*/
func RegisterModule(app *fiber.App) {

	userRepository := repositories.NewUserRepository(driver.DB)
	userActivationRepository := repositories.NewUserActionCodeRepository(driver.DB)
	aggregateRepository := repositories.NewRepositoryAggregate(userRepository, userActivationRepository)

	userActivationFactory := registerActivationFactory(userActivationRepository)

	userService := services.NewUserService(userRepository, userActivationRepository, aggregateRepository, userActivationFactory)
	userHandler := handlers.NewUserHandler(userService)

	routesInit := ApiRoute{
		UserHandler: *userHandler,
	}

	routesInit.Route(app)
}
