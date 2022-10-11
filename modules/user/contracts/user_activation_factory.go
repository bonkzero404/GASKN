package contracts

import (
	"go-starterkit-project/database/stores"
)

type UserActivationServiceFactory interface {
	CreateUserActivation(user *stores.User) (*stores.UserActivation, error)
}
