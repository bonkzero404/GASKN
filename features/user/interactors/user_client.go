package interactors

import (
	"github.com/bonkzero404/gaskn/database/stores"
	"github.com/bonkzero404/gaskn/features/user/dto"
)

type UserClient interface {
	CreateUserInvitation(clientId string, req *dto.UserInvitationRequest, invitedByUser string) (*dto.UserInvitationResponse, error)

	UserInviteAcceptance(clientId string, userId string, code string, accept stores.StatusInvitationType) (*dto.UserInvitationResponse, error)
}
