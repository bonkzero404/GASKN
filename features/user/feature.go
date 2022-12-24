package user

import (
	"github.com/bonkzero404/gaskn/driver"
	implementsRoleRepo "github.com/bonkzero404/gaskn/features/role/repositories/implements"
	implementsRoleAssignInteract "github.com/bonkzero404/gaskn/features/role_assignment/interactors/implements"
	roleAssignRepo "github.com/bonkzero404/gaskn/features/role_assignment/repositories/implements"
	"github.com/bonkzero404/gaskn/features/user/factories"
	implementsFactory "github.com/bonkzero404/gaskn/features/user/factories/implements"
	implementsInteract "github.com/bonkzero404/gaskn/features/user/interactors/implements"
	"github.com/bonkzero404/gaskn/features/user/repositories"
	"github.com/bonkzero404/gaskn/features/user/repositories/implements"
	"github.com/gofiber/fiber/v2"

	"github.com/bonkzero404/gaskn/config"
	"github.com/bonkzero404/gaskn/features/user/handlers"
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
	roleAssignmentRepository := roleAssignRepo.NewRoleAssignmentRepository(driver.DB)
	userActionFactory := registerActionCodeFactory(userActionCodeRepository)

	userService := implementsInteract.NewUser(
		userRepository,
		userActionCodeRepository,
		aggregateRepository,
		userActionFactory,
		userInvitationRepository,
	)

	userHandler := handlers.NewUserHandler(userService)

	repoUserRole := implementsRoleRepo.NewRoleRepository(driver.DB)
	repoUserRoleClient := implementsRoleRepo.NewRoleClientRepository(driver.DB)
	interactAssign := implementsRoleAssignInteract.NewRoleAssignment(repoUserRoleClient, repoUserRole, roleAssignmentRepository)

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

	var routesInitTenant = ApiRoute{}

	routesInit := ApiRoute{
		UserHandler:       *userHandler,
		UserClientHandler: *userClientHandler,
	}

	routesInit.Route(app)

	/////////////////////////
	// If tenant is enabled
	/////////////////////////
	if config.Config("TENANCY") == "true" {
		routesInitTenant = ApiRoute{
			UserHandler:       *userHandler,
			UserClientHandler: *userClientHandler,
		}
		routesInitTenant.RouteClient(app)
	}
}
