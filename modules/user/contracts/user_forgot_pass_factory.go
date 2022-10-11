package contracts

import "gaskn/database/stores"

type UserForgotPassServiceFactory interface {
	CreateUserForgotPass(user *stores.User) (*stores.UserActivation, error)
}
