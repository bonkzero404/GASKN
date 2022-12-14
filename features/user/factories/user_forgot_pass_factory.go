package factories

import "github.com/bonkzero404/gaskn/database/stores"

type UserForgotPassServiceFactory interface {
	CreateUserForgotPass(user *stores.User) (*stores.UserActionCode, error)
}
