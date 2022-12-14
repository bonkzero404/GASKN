package factories

import "gaskn/database/stores"

type ActionFactory interface {
	CreateActivation(user *stores.User) (*stores.UserActionCode, error)

	CreateForgotPassword(user *stores.User) (*stores.UserActionCode, error)

	CreateInvitation(user *stores.User, UrlFrontEndInvitation string, invitedBy string, role string, clientId string) (*stores.UserActionCode, error)
}
