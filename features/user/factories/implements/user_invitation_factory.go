package implements

import (
	"github.com/bonkzero404/gaskn/app/http"
	"github.com/bonkzero404/gaskn/app/mailer"
	"github.com/bonkzero404/gaskn/app/utils"
	"github.com/bonkzero404/gaskn/features/user/factories"
	"github.com/bonkzero404/gaskn/features/user/repositories"

	"github.com/gofiber/fiber/v2"

	"github.com/bonkzero404/gaskn/database/stores"
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
		return nil, &http.SetApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    err.Error(),
		}
	}

	sendMail := mailer.Mail{
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

	mailer.SendMail(&sendMail)

	return &userInvitation, nil
}
