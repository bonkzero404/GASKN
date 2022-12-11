package implements

import (
	"gaskn/features/user/factories"
	"gaskn/features/user/repositories"

	"github.com/gofiber/fiber/v2"

	"gaskn/database/stores"
	responseDto "gaskn/dto"
	"gaskn/utils"
)

type UserInvitationServiceFactory struct {
	UserInvitationRepository repositories.UserActionCodeRepository
}

func NewUserInvitationServiceFactory(
	UserInvitationRepository repositories.UserActionCodeRepository,
) factories.UserInvitationServiceFactory {
	return &UserInvitationServiceFactory{
		UserInvitationRepository: UserInvitationRepository,
	}
}

func (service UserInvitationServiceFactory) CreateUserInvitation(user *stores.User, urlInvitation string, invitedBy string) (*stores.UserActionCode, error) {
	codeGen := utils.StringWithCharset(32)

	userInvitation := stores.UserActionCode{
		UserId:  user.ID,
		Code:    codeGen,
		ActType: stores.INVITATION_CODE,
	}

	if err := service.UserInvitationRepository.CreateUserActionCode(&userInvitation).Error; err != nil {
		return &stores.UserActionCode{}, &responseDto.ApiErrorResponse{
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
		},
	}

	utils.SendMail(&sendMail)

	return &userInvitation, nil
}
