package services

import (
	"errors"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"gaskn/database/stores"
	respModel "gaskn/dto"
	"gaskn/modules/user/contracts"
	"gaskn/modules/user/dto"
	"gaskn/modules/user/services/factories"
	"gaskn/utils"
)

type UserService struct {
	UserRepository           contracts.UserRepository
	UserActionCodeRepository contracts.UserActionCodeRepository
	RepositoryAggregate      contracts.RepositoryAggregate
	ActionFactory            factories.ActionFactoryInterface
}

func NewUserService(
	UserRepository contracts.UserRepository,
	UserActivationRepository contracts.UserActionCodeRepository,
	RepositoryAggregate contracts.RepositoryAggregate,
	Factory factories.ActionFactoryInterface,
) contracts.UserService {
	return &UserService{
		UserRepository:           UserRepository,
		UserActionCodeRepository: UserActivationRepository,
		RepositoryAggregate:      RepositoryAggregate,
		ActionFactory:            Factory,
	}
}

func (service UserService) CreateUser(c *fiber.Ctx, user *dto.UserCreateRequest) (*dto.UserCreateResponse, error) {
	hashPassword, _ := utils.HashPassword(user.Password)

	userData := stores.User{
		FullName: user.FullName,
		Email:    user.Email,
		Phone:    user.Phone,
		Password: hashPassword,
	}

	activationCode := utils.StringWithCharset(32)

	userAvtivate := stores.UserActionCode{
		Code:    activationCode,
		ActType: stores.ACTIVATION_CODE,
	}

	result, err := service.RepositoryAggregate.CreateUser(&userData, &userAvtivate)

	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return &dto.UserCreateResponse{}, &respModel.ApiErrorResponse{
				StatusCode: fiber.StatusUnprocessableEntity,
				Message:    utils.Lang(c, "user:err:register-failed"),
			}
		}

		return &dto.UserCreateResponse{}, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    utils.Lang(c, "global:err:failed-unknown"),
		}
	}

	sendMail := respModel.Mail{
		To:           []string{user.Email},
		Subject:      "User Activation",
		TemplateHtml: "user_activation.html",
		BodyParam: map[string]interface{}{
			"Name": user.FullName,
			"Code": activationCode,
		},
	}

	utils.SendMail(&sendMail)

	response := dto.UserCreateResponse{
		ID:       userData.ID.String(),
		FullName: result.FullName,
		Email:    result.Email,
		Phone:    result.Phone,
		IsActive: userData.IsActive,
	}

	return &response, nil
}

func (service UserService) UserActivation(c *fiber.Ctx, email string, code string) (*dto.UserCreateResponse, error) {
	var user stores.User
	var userAct stores.UserActionCode

	errUser := service.UserRepository.FindUserByEmail(&user, email).Error

	if errors.Is(errUser, gorm.ErrRecordNotFound) {
		return &dto.UserCreateResponse{}, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusNotFound,
			Message:    utils.Lang(c, "user:err:user-not-found"),
		}
	}

	if user.IsActive {
		return &dto.UserCreateResponse{}, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    utils.Lang(c, "user:err:activate-already-active"),
		}
	}

	errAct := service.UserActionCodeRepository.FindUserActionCode(&userAct, user.ID.String(), code).Error

	if errors.Is(errAct, gorm.ErrRecordNotFound) {
		return &dto.UserCreateResponse{}, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusNotFound,
			Message:    utils.Lang(c, "user:err:activation-not-found"),
		}
	}

	t := time.Now()

	if userAct.ExpiredAt.Before(t) {
		return &dto.UserCreateResponse{}, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusGone,
			Message:    utils.Lang(c, "user:err:activation-expired"),
		}
	}

	userNew, errUserNew := service.RepositoryAggregate.UpdateUserActivation(user.ID.String(), true)

	if errUserNew != nil {
		return &dto.UserCreateResponse{}, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    errUserNew.Error(),
		}
	}

	service.RepositoryAggregate.UpdateActionCodeUsed(user.ID.String(), code)

	response := dto.UserCreateResponse{
		ID:       userNew.ID.String(),
		FullName: userNew.FullName,
		Email:    userNew.Email,
		Phone:    userNew.Phone,
		IsActive: userNew.IsActive,
	}

	return &response, nil
}

func (service UserService) CreateUserActivation(c *fiber.Ctx, email string, actType stores.ActCodeType) (map[string]interface{}, error) {
	var user stores.User

	errUser := service.UserRepository.FindUserByEmail(&user, email).Error

	if errors.Is(errUser, gorm.ErrRecordNotFound) {
		return nil, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusNotFound,
			Message:    utils.Lang(c, "user:err:user-not-found"),
		}
	}

	if user.IsActive && actType == stores.ACTIVATION_CODE {
		return nil, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    utils.Lang(c, "user:err:activate-already-active"),
		}
	}

	_, errActFactory := service.ActionFactory.Create(actType, &user)

	if errActFactory != nil {
		return nil, errActFactory
	}

	return map[string]interface{}{}, nil
}

func (service UserService) UpdatePassword(c *fiber.Ctx, forgotPassReq *dto.UserForgotPassActRequest) (map[string]interface{}, error) {
	var user stores.User
	var userAct stores.UserActionCode

	if forgotPassReq.Password != forgotPassReq.RepeatPassword {
		return nil, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    utils.Lang(c, "user:err:pass-match"),
		}
	}

	errUser := service.UserRepository.FindUserByEmail(&user, forgotPassReq.Email).Error

	if errors.Is(errUser, gorm.ErrRecordNotFound) {
		return nil, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusNotFound,
			Message:    utils.Lang(c, "user:err:user-not-found"),
		}
	}

	errAct := service.UserActionCodeRepository.FindUserActionCode(&userAct, user.ID.String(), forgotPassReq.Code).Error

	if errors.Is(errAct, gorm.ErrRecordNotFound) {
		return nil, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusNotFound,
			Message:    utils.Lang(c, "user:err:activation-not-found"),
		}
	}

	if userAct.IsUsed {
		return nil, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    utils.Lang(c, "user:err:pass-code-used"),
		}
	}

	t := time.Now()

	if userAct.ExpiredAt.Before(t) {
		return nil, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusGone,
			Message:    utils.Lang(c, "user:err:activation-expired"),
		}
	}

	go func() {
		hashPassword, _ := utils.HashPassword(user.Password)

		userData := stores.User{
			FullName: user.FullName,
			Email:    user.Email,
			Phone:    user.Phone,
			Password: hashPassword,
		}

		service.UserRepository.UpdatePassword(&userData)
		service.RepositoryAggregate.UpdateActionCodeUsed(user.ID.String(), forgotPassReq.Code)
	}()

	return map[string]interface{}{}, nil
}
