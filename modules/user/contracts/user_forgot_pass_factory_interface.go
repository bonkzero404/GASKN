package contracts

import "go-starterkit-project/database/stores"

type UserForgotPassServiceFactoryInterface interface {
	CreateUserForgotPass(user *stores.User) (*stores.UserActivation, error)
}
