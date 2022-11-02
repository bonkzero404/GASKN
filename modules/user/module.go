package user

import (
	"github.com/gofiber/fiber/v2"

	"gaskn/config"
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
	userInvitationRepository := repositories.NewUserInvitationRepository(driver.DB)
	userActionFactory := registerActionCodeFactory(userActionCodeRepository)

	userService := services.NewUserService(userRepository, userActionCodeRepository, aggregateRepository, userActionFactory)
	userHandler := handlers.NewUserHandler(userService)

	userClientService := services.NewUserClientService(userRepository, userActionCodeRepository, userInvitationRepository, aggregateRepository, userActionFactory)
	userClientHandler := handlers.NewUserClientHandler(userClientService)

	var routesInitTenant = ApiRouteClient{}

	routesInit := ApiRoute{
		UserHandler: *userHandler,
	}

	routesInit.Route(app)

	/////////////////////////
	// If tenant is enabled
	/////////////////////////
	if config.Config("TENANCY") == "true" {
		routesInitTenant = ApiRouteClient{
			UserClientHandler: *userClientHandler,
		}
		routesInitTenant.Route(app)
	}
}
