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

type UserForgotPassFactory struct {
	UserForgotPassRepository repositories.UserActionCodeRepository
}

func NewUserForgotPassFactory(UserForgotPassRepository repositories.UserActionCodeRepository) factories.UserForgotPassServiceFactory {
	return &UserForgotPassFactory{
		UserForgotPassRepository: UserForgotPassRepository,
	}
}

func (service UserForgotPassFactory) CreateUserForgotPass(user *stores.User) (*stores.UserActionCode, error) {
	codeGen := utils.StringWithCharset(32)

	userActivate := stores.UserActionCode{
		UserId:  user.ID,
		Code:    codeGen,
		ActType: stores.FORGOT_PASSWORD,
	}

	if err := service.UserForgotPassRepository.CreateUserActionCode(&userActivate).Error; err != nil {
		return nil, &http.SetApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    err.Error(),
		}
	}

	sendMail := mailer.Mail{
		To:           []string{user.Email},
		Subject:      "Forgot Password",
		TemplateHtml: "user_forgot_password.html",
		BodyParam: map[string]interface{}{
			"Name": user.FullName,
			"Code": codeGen,
		},
	}

	mailer.SendMail(&sendMail)

	return &userActivate, nil
}
