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
	Create(actionType stores.ActCodeType, user *stores.User) (*stores.UserActionCode, error)
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

func (factory ActionFactory) Create(actionType stores.ActCodeType, user *stores.User) (*stores.UserActionCode, error) {
	// Activation code
	if actionType == stores.ACTIVATION_CODE {
		userAct, err := factory.UserActivationServiceFactory.CreateUserActivation(user)

		if err != nil {
			return nil, err
		}

		return userAct, nil
	}

	// Forgot password
	if actionType == stores.FORGOT_PASSWORD {
		userAct, err := factory.UserForgotPassServiceFactory.CreateUserForgotPass(user)

		if err != nil {
			return nil, err
		}

		return userAct, nil
	}

	// User Invitation
	userAct, err := factory.UserInvitationServiceFactory.CreateUserInvitation(user)

	if err != nil {
		return nil, err
	}

	return userAct, nil
}
