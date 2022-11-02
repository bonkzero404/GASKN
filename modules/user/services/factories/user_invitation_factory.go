package factories

import (
	"github.com/gofiber/fiber/v2"

	"gaskn/database/stores"
	respModel "gaskn/dto"
	"gaskn/modules/user/contracts"
	"gaskn/utils"
)

type UserInvitationServiceFactory struct {
	UserInvitationRepository contracts.UserActionCodeRepository
}

func NewUserInvitationServiceFactory(
	UserInvitationRepository contracts.UserActionCodeRepository,
) contracts.UserInvitationServiceFactory {
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
		return &stores.UserActionCode{}, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    err.Error(),
		}
	}

	sendMail := respModel.Mail{
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
