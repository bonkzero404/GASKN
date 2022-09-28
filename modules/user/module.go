package user

import (
	"go-starterkit-project/database/driver"
	"go-starterkit-project/modules/user/domain/interfaces"
	"go-starterkit-project/modules/user/handlers"
	"go-starterkit-project/modules/user/repositories"
	"go-starterkit-project/modules/user/services"
	"go-starterkit-project/modules/user/services/factories"

	"github.com/gofiber/fiber/v2"
)

/*
*
Service factory registration
*/
func registerActivationFactory(userActivationRepository interfaces.UserActivationRepositoryInterface) factories.ActionFactoryInterface {
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
	userActivationRepository := repositories.NewUserActivationRepository(driver.DB)
	aggregateRepository := repositories.NewRepositoryAggregate(userRepository, userActivationRepository)

	userActivationFactory := registerActivationFactory(userActivationRepository)

	userService := services.NewUserService(userRepository, userActivationRepository, aggregateRepository, userActivationFactory)
	userHandler := handlers.NewUserHandler(userService)

	routesInit := ApiRoute{
		UserHandler: *userHandler,
	}

	routesInit.Route(app)
}
