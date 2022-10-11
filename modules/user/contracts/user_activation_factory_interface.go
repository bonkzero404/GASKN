package contracts

import (
	"go-starterkit-project/database/stores"
)

type UserActivationServiceFactoryInterface interface {
	CreateUserActivation(user *stores.User) (*stores.UserActivation, error)
}
