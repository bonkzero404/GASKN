package factories

import (
	"gaskn/database/stores"
	"gaskn/modules/user/contracts"
)

type ActionFactory struct {
	UserActivationServiceFactory contracts.UserActivationServiceFactory
	UserForgotPassServiceFactory contracts.UserForgotPassServiceFactory
	UserInvitationServiceFactory contracts.UserInvitationServiceFactory
}

type ActionFactoryInterface interface {
	CreateActivation(user *stores.User) (*stores.UserActionCode, error)
	CreateForgotPassword(user *stores.User) (*stores.UserActionCode, error)
	CreateInvitation(user *stores.User, UrlFrontEndInvitation string, invitedBy string) (*stores.UserActionCode, error)
}

func NewActionFactory(
	UserActivationServiceFactory contracts.UserActivationServiceFactory,
	UserForgotPassServiceFactory contracts.UserForgotPassServiceFactory,
	UserInvitationServiceFactory contracts.UserInvitationServiceFactory,
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

func (factory ActionFactory) CreateInvitation(user *stores.User, urlFrontEndInvitation string, invitedBy string) (*stores.UserActionCode, error) {
	userAct, err := factory.UserInvitationServiceFactory.CreateUserInvitation(user, urlFrontEndInvitation, invitedBy)

	if err != nil {
		return nil, err
	}

	return userAct, nil
}
