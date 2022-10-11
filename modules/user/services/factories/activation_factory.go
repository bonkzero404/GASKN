package factories

import (
	"gaskn/database/stores"
	"gaskn/modules/user/contracts"
)

type ActionFactory struct {
	UserActivationServiceFactory contracts.UserActivationServiceFactory
	UserForgotPassServiceFactory contracts.UserForgotPassServiceFactory
}

type ActionFactoryInterface interface {
	Create(actionType stores.ActivationType, user *stores.User) (*stores.UserActivation, error)
}

func NewActionFactory(
	userActivationServiceFactory contracts.UserActivationServiceFactory,
	userForgotPassServiceFactory contracts.UserForgotPassServiceFactory,
) ActionFactoryInterface {
	return &ActionFactory{
		UserActivationServiceFactory: userActivationServiceFactory,
		UserForgotPassServiceFactory: userForgotPassServiceFactory,
	}
}

func (factory ActionFactory) Create(actionType stores.ActivationType, user *stores.User) (*stores.UserActivation, error) {

	if actionType == stores.ACTIVATION_CODE {
		userAct, err := factory.UserActivationServiceFactory.CreateUserActivation(user)

		if err != nil {
			return nil, err
		}

		return userAct, nil
	}

	userAct, err := factory.UserForgotPassServiceFactory.CreateUserForgotPass(user)

	if err != nil {
		return nil, err
	}

	return userAct, nil
}
