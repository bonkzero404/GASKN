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

type UserActivationFactory struct {
	UserActivationRepository repositories.UserActionCodeRepository
}

func NewUserActivationFactory(
	UserActivationRepository repositories.UserActionCodeRepository,
) factories.UserActivationServiceFactory {
	return &UserActivationFactory{
		UserActivationRepository: UserActivationRepository,
	}
}

func (service UserActivationFactory) CreateUserActivation(user *stores.User) (*stores.UserActionCode, error) {
	codeGen := utils.StringWithCharset(32)

	userActivate := stores.UserActionCode{
		UserId:  user.ID,
		Code:    codeGen,
		ActType: stores.ACTIVATION_CODE,
	}

	if err := service.UserActivationRepository.CreateUserActionCode(&userActivate).Error; err != nil {
		return nil, &http.SetApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    err.Error(),
		}
	}

	sendMail := mailer.Mail{
		To:           []string{user.Email},
		Subject:      "User Activation",
		TemplateHtml: "user_activation.html",
		BodyParam: map[string]interface{}{
			"Name": user.FullName,
			"Code": codeGen,
		},
	}

	mailer.SendMail(&sendMail)

	return &userActivate, nil
}
