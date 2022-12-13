package implements

import (
	"gaskn/database/stores"
	factories2 "gaskn/features/user/factories"
)

type ActionFactory struct {
	UserActivationServiceFactory factories2.UserActivationServiceFactory
	UserForgotPassServiceFactory factories2.UserForgotPassServiceFactory
	UserInvitationServiceFactory factories2.UserInvitationServiceFactory
}

type ActionFactoryInterface interface {
	CreateActivation(user *stores.User) (*stores.UserActionCode, error)
	CreateForgotPassword(user *stores.User) (*stores.UserActionCode, error)
	CreateInvitation(user *stores.User, UrlFrontEndInvitation string, invitedBy string, role string, clientId string) (*stores.UserActionCode, error)
}

func NewActionFactory(
	UserActivationServiceFactory factories2.UserActivationServiceFactory,
	UserForgotPassServiceFactory factories2.UserForgotPassServiceFactory,
	UserInvitationServiceFactory factories2.UserInvitationServiceFactory,
) ActionFactoryInterface {
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
