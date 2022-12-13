package user

import (
	"gaskn/driver"
	implements5 "gaskn/features/role/repositories/implements"
	implements6 "gaskn/features/role_assignment/interactors/implements"
	implements4 "gaskn/features/user/factories/implements"
	implements2 "gaskn/features/user/interactors/implements"
	"gaskn/features/user/repositories"
	"gaskn/features/user/repositories/implements"
	"github.com/gofiber/fiber/v2"

	"gaskn/config"
	"gaskn/features/user/handlers"
)

/*
*
Service factory registration
*/
func registerActionCodeFactory(userActionCodeRepository repositories.UserActionCodeRepository) implements4.ActionFactoryInterface {
	actFactory := implements4.NewUserActivationServiceFactory(userActionCodeRepository)
	forgotPassFactory := implements4.NewUserForgotPassServiceFactory(userActionCodeRepository)
	userInvitationFactory := implements4.NewUserInvitationServiceFactory(userActionCodeRepository)

	return implements4.NewActionFactory(actFactory, forgotPassFactory, userInvitationFactory)
}

/*
*
This function is for registering repository - service - handler
*/
func RegisterFeature(app *fiber.App) {

	userRepository := implements.NewUserRepository(driver.DB)
	userActionCodeRepository := implements.NewUserActionCodeRepository(driver.DB)
	aggregateRepository := implements.NewRepositoryAggregate(userRepository, userActionCodeRepository)
	userInvitationRepository := implements.NewUserInvitationRepository(driver.DB)
	userActionFactory := registerActionCodeFactory(userActionCodeRepository)

	userService := implements2.NewUser(userRepository, userActionCodeRepository, aggregateRepository, userActionFactory)
	userHandler := handlers.NewUserHandler(userService)

	repoUserRole := implements5.NewRoleRepository(driver.DB)
	repoUserRoleClient := implements5.NewRoleClientRepository(driver.DB)
	interactAssign := implements6.NewRoleAssignment(repoUserRoleClient, repoUserRole)

	userClientService := implements2.NewUserClient(
		userRepository,
		userActionCodeRepository,
		userInvitationRepository,
		aggregateRepository,
		userActionFactory,
		repoUserRoleClient,
		interactAssign,
	)
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
