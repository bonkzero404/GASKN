package contracts

import (
	"go-starterkit-project/database/stores"
)

type RepositoryAggregate interface {
	CreateUser(user *stores.User, userActivate *stores.UserActivation) (*stores.User, error)

	UpdateUserActivation(id string, stat bool) (*stores.User, error)

	UpdatePassword(id string, password string) (*stores.User, error)

	UpdateActivationCodeUsed(userId string, code string) (*stores.UserActivation, error)
}
