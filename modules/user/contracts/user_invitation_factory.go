package contracts

import "gaskn/database/stores"

type UserInvitationServiceFactory interface {
	CreateUserInvitation(user *stores.User, urlInvitation string, invitedBy string) (*stores.UserActionCode, error)
}
