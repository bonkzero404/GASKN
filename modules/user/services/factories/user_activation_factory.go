package factories

import (
	"go-starterkit-project/database/stores"
	respModel "go-starterkit-project/dto"
	"go-starterkit-project/modules/user/contracts"
	"go-starterkit-project/utils"

	"github.com/gofiber/fiber/v2"
)

type UserActivationServiceFactory struct {
	UserActivationRepository contracts.UserActivationRepository
}

func NewUserActivationServiceFactory(
	userActivationRepository contracts.UserActivationRepository,
) contracts.UserActivationServiceFactory {
	return &UserActivationServiceFactory{
		UserActivationRepository: userActivationRepository,
	}
}

func (service UserActivationServiceFactory) CreateUserActivation(user *stores.User) (*stores.UserActivation, error) {
	codeGen := utils.StringWithCharset(32)

	userActivate := stores.UserActivation{
		UserId:  user.ID,
		Code:    codeGen,
		ActType: stores.ACTIVATION_CODE,
	}

	if err := service.UserActivationRepository.CreateUserActivation(&userActivate).Error; err != nil {
		return &stores.UserActivation{}, &respModel.ApiErrorResponse{
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
