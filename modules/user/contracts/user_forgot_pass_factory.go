package contracts

import "go-starterkit-project/database/stores"

type UserForgotPassServiceFactory interface {
	CreateUserForgotPass(user *stores.User) (*stores.UserActivation, error)
}
