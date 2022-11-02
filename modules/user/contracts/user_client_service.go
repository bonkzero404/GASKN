package contracts

import (
	"github.com/gofiber/fiber/v2"

	"gaskn/database/stores"
	"gaskn/modules/user/dto"
)

type UserClientService interface {
	CreateUserInvitation(c *fiber.Ctx, req *dto.UserInvitationRequest, invitedByUser string, actType stores.ActCodeType) (map[string]interface{}, error)

	// InvitationApproved(c *fiber.Ctx, forgotPassReq *dto.UserForgotPassActRequest) (map[string]interface{}, error)
}
