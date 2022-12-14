package interactors

import (
	"github.com/gofiber/fiber/v2"

	"github.com/bonkzero404/gaskn/database/stores"
	"github.com/bonkzero404/gaskn/features/user/dto"
)

type UserClient interface {
	CreateUserInvitation(c *fiber.Ctx, req *dto.UserInvitationRequest, invitedByUser string) (*dto.UserInvitationResponse, error)

	UserInviteAcceptance(c *fiber.Ctx, code string, accept stores.StatusInvitationType) (*dto.UserInvitationResponse, error)
}
