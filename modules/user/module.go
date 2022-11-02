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
func registerActionCodeFactory(userActionCodeRepository contracts.UserActionCodeRepository) factories.ActionFactoryInterface {
	actFactory := factories.NewUserActivationServiceFactory(userActionCodeRepository)
	forgotPassFactory := factories.NewUserForgotPassServiceFactory(userActionCodeRepository)
	userInvitationFactory := factories.NewUserInvitationServiceFactory(userActionCodeRepository)

	return factories.NewActionFactory(actFactory, forgotPassFactory, userInvitationFactory)
}

/*
*
This function is for registering repository - service - handler
*/
func RegisterModule(app *fiber.App) {

	userRepository := repositories.NewUserRepository(driver.DB)
	userActionCodeRepository := repositories.NewUserActionCodeRepository(driver.DB)
	aggregateRepository := repositories.NewRepositoryAggregate(userRepository, userActionCodeRepository)

	userActivationFactory := registerActionCodeFactory(userActionCodeRepository)

	userService := services.NewUserService(userRepository, userActionCodeRepository, aggregateRepository, userActivationFactory)
	userHandler := handlers.NewUserHandler(userService)

	routesInit := ApiRoute{
		UserHandler: *userHandler,
	}

	routesInit.Route(app)
}
