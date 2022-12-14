package factories

import (
	"github.com/bonkzero404/gaskn/database/stores"
)

type UserInvitationServiceFactory interface {
	CreateUserInvitation(user *stores.User, urlInvitation string, invitedBy string, role string, clientId string) (*stores.UserActionCode, error)
}
