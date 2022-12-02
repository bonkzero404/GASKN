package factories

import (
	"github.com/gofiber/fiber/v2"

	"gaskn/database/stores"
	respModel "gaskn/dto"
	"gaskn/features/user/contracts"
	"gaskn/utils"
)

type UserActivationServiceFactory struct {
	UserActivationRepository contracts.UserActionCodeRepository
}

func NewUserActivationServiceFactory(
	UserActivationRepository contracts.UserActionCodeRepository,
) contracts.UserActivationServiceFactory {
	return &UserActivationServiceFactory{
		UserActivationRepository: UserActivationRepository,
	}
}

func (service UserActivationServiceFactory) CreateUserActivation(user *stores.User) (*stores.UserActionCode, error) {
	codeGen := utils.StringWithCharset(32)

	userActivate := stores.UserActionCode{
		UserId:  user.ID,
		Code:    codeGen,
		ActType: stores.ACTIVATION_CODE,
	}

	if err := service.UserActivationRepository.CreateUserActionCode(&userActivate).Error; err != nil {
		return &stores.UserActionCode{}, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    err.Error(),
		}
	}

	sendMail := respModel.Mail{
		To:           []string{user.Email},
		Subject:      "User Activation",
		TemplateHtml: "user_activation.html",
		BodyParam: map[string]interface{}{
			"Name": user.FullName,
			"Code": codeGen,
		},
	}

	utils.SendMail(&sendMail)

	return &userActivate, nil
}
