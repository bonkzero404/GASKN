package contracts

import (
	"gaskn/database/stores"
)

type UserActivationServiceFactory interface {
	CreateUserActivation(user *stores.User) (*stores.UserActionCode, error)
}
