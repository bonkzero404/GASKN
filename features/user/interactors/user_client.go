package interactors

import (
	"github.com/gofiber/fiber/v2"

	"gaskn/database/stores"
	"gaskn/features/user/dto"
)

type UserClient interface {
	CreateUserInvitation(c *fiber.Ctx, req *dto.UserInvitationRequest, invitedByUser string) (map[string]interface{}, error)

	UserInviteAcceptance(c *fiber.Ctx, code string, accept stores.StatusInvitationType) (*stores.UserInvitation, error)
}
