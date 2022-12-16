package implements

import (
	"github.com/bonkzero404/gaskn/features/user/factories"
	"github.com/bonkzero404/gaskn/features/user/repositories"

	"github.com/gofiber/fiber/v2"

	"github.com/bonkzero404/gaskn/database/stores"
	responseDto "github.com/bonkzero404/gaskn/dto"
	"github.com/bonkzero404/gaskn/utils"
)

type UserInvitationFactory struct {
	UserInvitationRepository repositories.UserActionCodeRepository
}

func NewUserInvitationFactory(
	UserInvitationRepository repositories.UserActionCodeRepository,
) factories.UserInvitationServiceFactory {
	return &UserInvitationFactory{
		UserInvitationRepository: UserInvitationRepository,
	}
}

func (service UserInvitationFactory) CreateUserInvitation(user *stores.User, urlInvitation string, invitedBy string, role string, clientId string) (*stores.UserActionCode, error) {
	codeGen := utils.StringWithCharset(32)

	userInvitation := stores.UserActionCode{
		UserId:  user.ID,
		Code:    codeGen,
		ActType: stores.INVITATION_CODE,
	}

	if err := service.UserInvitationRepository.CreateUserActionCode(&userInvitation).Error; err != nil {
		return nil, &responseDto.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    err.Error(),
		}
	}

	sendMail := responseDto.Mail{
		To:           []string{user.Email},
		Subject:      "User Invitation",
		TemplateHtml: "user_invitation.html",
		BodyParam: map[string]interface{}{
			"Name":          user.FullName,
			"Code":          codeGen,
			"UrlInvitation": urlInvitation,
			"InvitedBy":     invitedBy,
			"Role":          role,
			"ClientId":      clientId,
		},
	}

	utils.SendMail(&sendMail)

	return &userInvitation, nil
}
