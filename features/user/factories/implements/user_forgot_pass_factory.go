package implements

import (
	"gaskn/features/user/factories"
	"gaskn/features/user/repositories"

	"github.com/gofiber/fiber/v2"

	"gaskn/database/stores"
	responseDto "gaskn/dto"
	"gaskn/utils"
)

type UserForgotPassServiceFactory struct {
	UserForgotPassRepository repositories.UserActionCodeRepository
}

func NewUserForgotPassServiceFactory(UserForgotPassRepository repositories.UserActionCodeRepository) factories.UserForgotPassServiceFactory {
	return &UserForgotPassServiceFactory{
		UserForgotPassRepository: UserForgotPassRepository,
	}
}

func (service UserForgotPassServiceFactory) CreateUserForgotPass(user *stores.User) (*stores.UserActionCode, error) {
	codeGen := utils.StringWithCharset(32)

	userActivate := stores.UserActionCode{
		UserId:  user.ID,
		Code:    codeGen,
		ActType: stores.FORGOT_PASSWORD,
	}

	if err := service.UserForgotPassRepository.CreateUserActionCode(&userActivate).Error; err != nil {
		return &stores.UserActionCode{}, &responseDto.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    err.Error(),
		}
	}

	sendMail := responseDto.Mail{
		To:           []string{user.Email},
		Subject:      "Forgot Password",
		TemplateHtml: "user_forgot_password.html",
		BodyParam: map[string]interface{}{
			"Name": user.FullName,
			"Code": codeGen,
		},
	}

	utils.SendMail(&sendMail)

	return &userActivate, nil
}
