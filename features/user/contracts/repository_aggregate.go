package contracts

import (
	"gaskn/database/stores"
)

type RepositoryAggregate interface {
	CreateUser(user *stores.User, userActivate *stores.UserActionCode) (*stores.User, error)

	UpdateUserActivation(id string, stat bool) (*stores.User, error)

	UpdatePassword(id string, password string) (*stores.User, error)

	UpdateActionCodeUsed(userId string, code string) (*stores.UserActionCode, error)
}
