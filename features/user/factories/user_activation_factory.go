package factories

import (
	"github.com/bonkzero404/gaskn/database/stores"
)

type UserActivationServiceFactory interface {
	CreateUserActivation(user *stores.User) (*stores.UserActionCode, error)
}
