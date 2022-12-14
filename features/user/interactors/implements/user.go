package implements

import (
	"errors"
	"github.com/bonkzero404/gaskn/features/user/factories"
	"github.com/bonkzero404/gaskn/features/user/interactors"
	"github.com/bonkzero404/gaskn/features/user/repositories"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/bonkzero404/gaskn/database/stores"
	responseDto "github.com/bonkzero404/gaskn/dto"
	"github.com/bonkzero404/gaskn/features/user/dto"
	"github.com/bonkzero404/gaskn/utils"
)

type User struct {
	UserRepository           repositories.UserRepository
	UserActionCodeRepository repositories.UserActionCodeRepository
	RepositoryAggregate      repositories.RepositoryAggregate
	ActionFactory            factories.ActionFactory
}

func NewUser(
	UserRepository repositories.UserRepository,
	UserActionCodeRepository repositories.UserActionCodeRepository,
	RepositoryAggregate repositories.RepositoryAggregate,
	Factory factories.ActionFactory,
) interactors.User {
	return &User{
		UserRepository:           UserRepository,
		UserActionCodeRepository: UserActionCodeRepository,
		RepositoryAggregate:      RepositoryAggregate,
		ActionFactory:            Factory,
	}
}

func (interact User) CreateUser(c *fiber.Ctx, user *dto.UserCreateRequest) (*dto.UserCreateResponse, error) {
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

	result, err := interact.RepositoryAggregate.CreateUser(&userData, &userAvtivate)

	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return &dto.UserCreateResponse{}, &responseDto.ApiErrorResponse{
				StatusCode: fiber.StatusUnprocessableEntity,
				Message:    utils.Lang(c, "user:err:register-failed"),
			}
		}

		return &dto.UserCreateResponse{}, &responseDto.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    utils.Lang(c, "global:err:failed-unknown"),
		}
	}

	sendMail := responseDto.Mail{
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

func (interact User) UserActivation(c *fiber.Ctx, code string) (*dto.UserCreateResponse, error) {
	var user stores.User
	var userAct stores.UserActionCode

	errAct := interact.UserActionCodeRepository.FindActionCode(&userAct, code).Error

	if errors.Is(errAct, gorm.ErrRecordNotFound) {
		return &dto.UserCreateResponse{}, &responseDto.ApiErrorResponse{
			StatusCode: fiber.StatusNotFound,
			Message:    utils.Lang(c, "user:err:activation-not-found"),
		}
	}

	errUser := interact.UserRepository.FindUserById(&user, userAct.UserId.String()).Error

	if errors.Is(errUser, gorm.ErrRecordNotFound) {
		return &dto.UserCreateResponse{}, &responseDto.ApiErrorResponse{
			StatusCode: fiber.StatusNotFound,
			Message:    utils.Lang(c, "user:err:user-not-found"),
		}
	}

	if user.IsActive {
		return &dto.UserCreateResponse{}, &responseDto.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    utils.Lang(c, "user:err:activate-already-active"),
		}
	}

	t := time.Now()

	if userAct.ExpiredAt.Before(t) {
		return &dto.UserCreateResponse{}, &responseDto.ApiErrorResponse{
			StatusCode: fiber.StatusGone,
			Message:    utils.Lang(c, "user:err:activation-expired"),
		}
	}

	userNew, errUserNew := interact.RepositoryAggregate.UpdateUserActivation(user.ID.String(), true)

	if errUserNew != nil {
		return &dto.UserCreateResponse{}, &responseDto.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    errUserNew.Error(),
		}
	}

	_, err := interact.RepositoryAggregate.UpdateActionCodeUsed(user.ID.String(), code)
	if err != nil {
		return nil, err
	}

	response := dto.UserCreateResponse{
		ID:       userNew.ID.String(),
		FullName: userNew.FullName,
		Email:    userNew.Email,
		Phone:    userNew.Phone,
		IsActive: userNew.IsActive,
	}

	return &response, nil
}

func (interact User) CreateUserActivation(c *fiber.Ctx, email string, actType stores.ActCodeType) (map[string]interface{}, error) {
	var user stores.User

	errUser := interact.UserRepository.FindUserByEmail(&user, email).Error

	if errors.Is(errUser, gorm.ErrRecordNotFound) {
		return nil, &responseDto.ApiErrorResponse{
			StatusCode: fiber.StatusNotFound,
			Message:    utils.Lang(c, "user:err:user-not-found"),
		}
	}

	if user.IsActive && actType == stores.ACTIVATION_CODE {
		return nil, &responseDto.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    utils.Lang(c, "user:err:activate-already-active"),
		}
	}

	_, errActFactory := interact.ActionFactory.CreateActivation(&user)

	if errActFactory != nil {
		return nil, errActFactory
	}

	return nil, nil
}

func (interact User) UpdatePassword(c *fiber.Ctx, forgotPassReq *dto.UserForgotPassActRequest) (map[string]interface{}, error) {
	var user stores.User
	var userAct stores.UserActionCode

	if forgotPassReq.Password != forgotPassReq.RepeatPassword {
		return nil, &responseDto.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    utils.Lang(c, "user:err:pass-match"),
		}
	}

	errUser := interact.UserRepository.FindUserByEmail(&user, forgotPassReq.Email).Error

	if errors.Is(errUser, gorm.ErrRecordNotFound) {
		return nil, &responseDto.ApiErrorResponse{
			StatusCode: fiber.StatusNotFound,
			Message:    utils.Lang(c, "user:err:user-not-found"),
		}
	}

	errAct := interact.UserActionCodeRepository.FindUserActionCode(&userAct, user.ID.String(), forgotPassReq.Code).Error

	if errors.Is(errAct, gorm.ErrRecordNotFound) {
		return nil, &responseDto.ApiErrorResponse{
			StatusCode: fiber.StatusNotFound,
			Message:    utils.Lang(c, "user:err:activation-not-found"),
		}
	}

	if userAct.IsUsed {
		return nil, &responseDto.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    utils.Lang(c, "user:err:pass-code-used"),
		}
	}

	t := time.Now()

	if userAct.ExpiredAt.Before(t) {
		return nil, &responseDto.ApiErrorResponse{
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

		interact.UserRepository.UpdatePassword(&userData)
		_, err := interact.RepositoryAggregate.UpdateActionCodeUsed(user.ID.String(), forgotPassReq.Code)
		if err != nil {
			return
		}
	}()

	return nil, nil
}
