package user

import (
	"gaskn/driver"
	"github.com/gofiber/fiber/v2"

	"gaskn/config"
	"gaskn/features/user/contracts"
	"gaskn/features/user/handlers"
	"gaskn/features/user/repositories"
	"gaskn/features/user/services"
	"gaskn/features/user/services/factories"
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
func RegisterFeature(app *fiber.App) {

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
		UserHandler:       *userHandler,
		UserClientHandler: *userClientHandler,
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
