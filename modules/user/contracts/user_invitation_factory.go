package contracts

import "gaskn/database/stores"

type UserInvitationServiceFactory interface {
	CreateUserInvitation(user *stores.User) (*stores.UserActionCode, error)
}
