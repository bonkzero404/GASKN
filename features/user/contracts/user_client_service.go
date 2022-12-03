package contracts

import (
	"github.com/gofiber/fiber/v2"

	"gaskn/database/stores"
	"gaskn/features/user/dto"
)

type UserClientService interface {
	CreateUserInvitation(c *fiber.Ctx, req *dto.UserInvitationRequest, invitedByUser string, actType stores.ActCodeType) (map[string]interface{}, error)

	UserInviteAcceptance(c *fiber.Ctx, code string, accept stores.StatusInvitationType) (*stores.UserInvitation, error)
}
