package user

import (
	"gaskn/driver"
	implementsRoleRepo "gaskn/features/role/repositories/implements"
	implementsRoleAssignInteract "gaskn/features/role_assignment/interactors/implements"
	"gaskn/features/user/factories"
	implementsFactory "gaskn/features/user/factories/implements"
	implementsInteract "gaskn/features/user/interactors/implements"
	"gaskn/features/user/repositories"
	"gaskn/features/user/repositories/implements"
	"github.com/gofiber/fiber/v2"

	"gaskn/config"
	"gaskn/features/user/handlers"
)

func registerActionCodeFactory(userActionCodeRepository repositories.UserActionCodeRepository) factories.ActionFactory {
	actFactory := implementsFactory.NewUserActivationFactory(userActionCodeRepository)
	forgotPassFactory := implementsFactory.NewUserForgotPassFactory(userActionCodeRepository)
	userInvitationFactory := implementsFactory.NewUserInvitationFactory(userActionCodeRepository)

	return implementsFactory.NewActionFactory(actFactory, forgotPassFactory, userInvitationFactory)
}

func RegisterFeature(app *fiber.App) {

	userRepository := implements.NewUserRepository(driver.DB)
	userActionCodeRepository := implements.NewUserActionCodeRepository(driver.DB)
	aggregateRepository := implements.NewRepositoryAggregate(userRepository, userActionCodeRepository)
	userInvitationRepository := implements.NewUserInvitationRepository(driver.DB)
	userActionFactory := registerActionCodeFactory(userActionCodeRepository)

	userService := implementsInteract.NewUser(userRepository, userActionCodeRepository, aggregateRepository, userActionFactory)
	userHandler := handlers.NewUserHandler(userService)

	repoUserRole := implementsRoleRepo.NewRoleRepository(driver.DB)
	repoUserRoleClient := implementsRoleRepo.NewRoleClientRepository(driver.DB)
	interactAssign := implementsRoleAssignInteract.NewRoleAssignment(repoUserRoleClient, repoUserRole)

	userClientService := implementsInteract.NewUserClient(
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
