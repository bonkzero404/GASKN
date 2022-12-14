package implements

import (
	"gaskn/database/stores"
	"gaskn/features/user/factories"
)

type ActionFactory struct {
	UserActivationServiceFactory factories.UserActivationServiceFactory
	UserForgotPassServiceFactory factories.UserForgotPassServiceFactory
	UserInvitationServiceFactory factories.UserInvitationServiceFactory
}

func NewActionFactory(
	UserActivationServiceFactory factories.UserActivationServiceFactory,
	UserForgotPassServiceFactory factories.UserForgotPassServiceFactory,
	UserInvitationServiceFactory factories.UserInvitationServiceFactory,
) factories.ActionFactory {
	return &ActionFactory{
		UserActivationServiceFactory: UserActivationServiceFactory,
		UserForgotPassServiceFactory: UserForgotPassServiceFactory,
		UserInvitationServiceFactory: UserInvitationServiceFactory,
	}
}

func (factory ActionFactory) CreateActivation(user *stores.User) (*stores.UserActionCode, error) {
	userAct, err := factory.UserActivationServiceFactory.CreateUserActivation(user)

	if err != nil {
		return nil, err
	}

	return userAct, nil
}

func (factory ActionFactory) CreateForgotPassword(user *stores.User) (*stores.UserActionCode, error) {
	userAct, err := factory.UserForgotPassServiceFactory.CreateUserForgotPass(user)

	if err != nil {
		return nil, err
	}

	return userAct, nil
}

func (factory ActionFactory) CreateInvitation(user *stores.User, urlFrontEndInvitation string, invitedBy string, role string, clientId string) (*stores.UserActionCode, error) {
	userAct, err := factory.UserInvitationServiceFactory.CreateUserInvitation(user, urlFrontEndInvitation, invitedBy, role, clientId)

	if err != nil {
		return nil, err
	}

	return userAct, nil
}
