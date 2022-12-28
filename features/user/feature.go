package user

import (
	"github.com/bonkzero404/gaskn/driver"
	roleRepo "github.com/bonkzero404/gaskn/features/role/repositories/implements"
	roleAssignInteract "github.com/bonkzero404/gaskn/features/role_assignment/interactors/implements"
	roleAssignRepo "github.com/bonkzero404/gaskn/features/role_assignment/repositories/implements"
	"github.com/bonkzero404/gaskn/features/user/factories"
	factory "github.com/bonkzero404/gaskn/features/user/factories/implements"
	userInteract "github.com/bonkzero404/gaskn/features/user/interactors/implements"
	"github.com/bonkzero404/gaskn/features/user/repositories"
	"github.com/bonkzero404/gaskn/features/user/repositories/implements"
	"github.com/gofiber/fiber/v2"

	"github.com/bonkzero404/gaskn/config"
	"github.com/bonkzero404/gaskn/features/user/handlers"
)

func registerActionCodeFactory(userActionCodeRepository repositories.UserActionCodeRepository) factories.ActionFactory {
	actFactory := factory.NewUserActivationFactory(userActionCodeRepository)
	forgotPassFactory := factory.NewUserForgotPassFactory(userActionCodeRepository)
	userInvitationFactory := factory.NewUserInvitationFactory(userActionCodeRepository)

	return factory.NewActionFactory(actFactory, forgotPassFactory, userInvitationFactory)
}

func RegisterFeature(app *fiber.App) {

	userRepository := implements.NewUserRepository(driver.DB)
	userActionCodeRepository := implements.NewUserActionCodeRepository(driver.DB)
	aggregateRepository := implements.NewRepositoryAggregate(userRepository, userActionCodeRepository)
	userInvitationRepository := implements.NewUserInvitationRepository(driver.DB)
	roleAssignmentRepository := roleAssignRepo.NewRoleAssignmentRepository(driver.DB)
	userActionFactory := registerActionCodeFactory(userActionCodeRepository)

	user := userInteract.NewUser(
		userRepository,
		userActionCodeRepository,
		aggregateRepository,
		userActionFactory,
		userInvitationRepository,
	)

	userHandler := handlers.NewUserHandler(user)

	repoUserRole := roleRepo.NewRoleRepository(driver.DB)
	repoUserRoleClient := roleRepo.NewRoleClientRepository(driver.DB)
	interactAssign := roleAssignInteract.NewRoleAssignment(repoUserRoleClient, repoUserRole, roleAssignmentRepository)

	userClient := userInteract.NewUserClient(
		userRepository,
		userActionCodeRepository,
		userInvitationRepository,
		aggregateRepository,
		userActionFactory,
		repoUserRoleClient,
		interactAssign,
	)
	userClientHandler := handlers.NewUserClientHandler(userClient)

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
