package factories

import (
	"github.com/gofiber/fiber/v2"

	"gaskn/database/stores"
	respModel "gaskn/dto"
	"gaskn/modules/user/contracts"
	"gaskn/utils"
)

type UserForgotPassServiceFactory struct {
	UserForgotPassRepository contracts.UserActionCodeRepository
}

func NewUserForgotPassServiceFactory(UserForgotPassRepository contracts.UserActionCodeRepository) contracts.UserForgotPassServiceFactory {
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
		return &stores.UserActionCode{}, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    err.Error(),
		}
	}

	sendMail := respModel.Mail{
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
