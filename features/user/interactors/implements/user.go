package implements

import (
	"errors"
	"github.com/bonkzero404/gaskn/app/facades"
	"github.com/bonkzero404/gaskn/app/mailer"
	"github.com/bonkzero404/gaskn/app/translation"
	"github.com/bonkzero404/gaskn/app/utils"
	"github.com/bonkzero404/gaskn/config"
	"github.com/bonkzero404/gaskn/features/user/factories"
	"github.com/bonkzero404/gaskn/features/user/interactors"
	"github.com/bonkzero404/gaskn/features/user/repositories"
	"github.com/google/uuid"
	"strings"
	"time"

	"gorm.io/gorm"

	"github.com/bonkzero404/gaskn/database/stores"
	"github.com/bonkzero404/gaskn/features/user/dto"
)

type User struct {
	UserRepository           repositories.UserRepository
	UserActionCodeRepository repositories.UserActionCodeRepository
	RepositoryAggregate      repositories.RepositoryAggregate
	ActionFactory            factories.ActionFactory
	UserInvitationRepository repositories.UserInvitationRepository
}

func NewUser(
	UserRepository repositories.UserRepository,
	UserActionCodeRepository repositories.UserActionCodeRepository,
	RepositoryAggregate repositories.RepositoryAggregate,
	Factory factories.ActionFactory,
	UserInvitationRepository repositories.UserInvitationRepository,
) interactors.User {
	return &User{
		UserRepository:           UserRepository,
		UserActionCodeRepository: UserActionCodeRepository,
		RepositoryAggregate:      RepositoryAggregate,
		ActionFactory:            Factory,
		UserInvitationRepository: UserInvitationRepository,
	}
}

func (repository User) CreateUser(clientId string, user *dto.UserCreateRequest, isInternalRegister bool) (*dto.UserCreateResponse, error) {
	hashPassword, _ := utils.HashPassword(user.Password)

	userData := stores.User{
		FullName: user.FullName,
		Email:    user.Email,
		Phone:    user.Phone,
		Password: hashPassword,
	}

	activationCode := utils.StringWithCharset(32)

	var userActionCode = stores.UserActionCode{}

	userActionCode = stores.UserActionCode{
		Code:    activationCode,
		ActType: stores.ACTIVATION_CODE,
	}

	// Check admin or client admin create a user
	if isInternalRegister {
		userActionCode.IsUsed = true
		userData.IsActive = true
	}

	result, err := repository.RepositoryAggregate.CreateUser(&userData, &userActionCode)

	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return nil, &facades.ResponseError{
				StatusCode: facades.AppErrUnprocessable,
				Message:    translation.Lang(config.UserErrRegister),
			}
		}

		return nil, &facades.ResponseError{
			StatusCode: facades.AppErrUnprocessable,
			Message:    translation.Lang(config.GlobalErrUnknown),
		}
	}

	if clientId != "" {
		cUuid, _ := uuid.Parse(clientId)
		clientAssign := stores.ClientAssignment{
			ClientId: cUuid,
			UserId:   userData.ID,
			IsActive: true,
		}

		repository.UserInvitationRepository.CreateClientAssignment(&clientAssign)
	}

	var sendMail = mailer.Mail{}

	if isInternalRegister {
		sendMail = mailer.Mail{
			To:           []string{user.Email},
			Subject:      "User Invitation",
			TemplateHtml: "user_creation.html",
			BodyParam: map[string]interface{}{
				"Name":     user.FullName,
				"Client":   config.Config("APP_NAME"),
				"Email":    user.Email,
				"Password": user.Password,
			},
		}
	} else {
		sendMail = mailer.Mail{
			To:           []string{user.Email},
			Subject:      "User Activation",
			TemplateHtml: "user_activation.html",
			BodyParam: map[string]interface{}{
				"Name": user.FullName,
				"Code": activationCode,
			},
		}
	}

	mailer.SendMail(&sendMail)

	response := dto.UserCreateResponse{
		ID:       userData.ID.String(),
		FullName: result.FullName,
		Email:    result.Email,
		Phone:    result.Phone,
		IsActive: userData.IsActive,
	}

	return &response, nil
}

func (repository User) UserActivation(code string) (*dto.UserCreateResponse, error) {
	var user stores.User
	var userAct stores.UserActionCode

	errAct := repository.UserActionCodeRepository.FindActionCode(&userAct, code).Error

	if errors.Is(errAct, gorm.ErrRecordNotFound) {
		return nil, &facades.ResponseError{
			StatusCode: facades.AppErrNotFound,
			Message:    translation.Lang(config.UserErrActivationNotFound),
		}
	}

	errUser := repository.UserRepository.FindUserById(&user, userAct.UserId.String()).Error

	if errors.Is(errUser, gorm.ErrRecordNotFound) {
		return nil, &facades.ResponseError{
			StatusCode: facades.AppErrNotFound,
			Message:    translation.Lang(config.UserErrNotFound),
		}
	}

	if user.IsActive {
		return nil, &facades.ResponseError{
			StatusCode: facades.AppErrUnprocessable,
			Message:    translation.Lang(config.UserErrAlreadyActive),
		}
	}

	t := time.Now()

	if userAct.ExpiredAt.Before(t) {
		return nil, &facades.ResponseError{
			StatusCode: facades.AppErrExpire,
			Message:    translation.Lang(config.UserErrActivationExpired),
		}
	}

	userNew, errUserNew := repository.RepositoryAggregate.UpdateUserActivation(user.ID.String(), true)

	if errUserNew != nil {
		return nil, &facades.ResponseError{
			StatusCode: facades.AppErrUnprocessable,
			Message:    errUserNew.Error(),
		}
	}

	_, err := repository.RepositoryAggregate.UpdateActionCodeUsed(user.ID.String(), code)
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

func (repository User) CreateUserAction(email string, actType stores.ActCodeType) (any, error) {
	var user stores.User

	errUser := repository.UserRepository.FindUserByEmail(&user, email).Error

	if errors.Is(errUser, gorm.ErrRecordNotFound) {
		return nil, &facades.ResponseError{
			StatusCode: facades.AppErrNotFound,
			Message:    translation.Lang(config.UserErrNotFound),
		}
	}

	if user.IsActive && actType == stores.ACTIVATION_CODE {
		return nil, &facades.ResponseError{
			StatusCode: facades.AppErrUnprocessable,
			Message:    translation.Lang(config.UserErrAlreadyActive),
		}
	}

	if actType == stores.FORGOT_PASSWORD {
		_, errActFactory := repository.ActionFactory.CreateForgotPassword(&user)
		if errActFactory != nil {
			return nil, errActFactory
		}

		return nil, nil
	}
	_, errActFactory := repository.ActionFactory.CreateActivation(&user)

	if errActFactory != nil {
		return nil, errActFactory
	}

	return nil, nil
}

func (repository User) UpdatePassword(forgotPassReq *dto.UserForgotPassActRequest) (any, error) {
	var user stores.User
	var userAct stores.UserActionCode

	if forgotPassReq.Password != forgotPassReq.RepeatPassword {
		return nil, &facades.ResponseError{
			StatusCode: facades.AppErrUnprocessable,
			Message:    translation.Lang(config.UserErrPassMatch),
		}
	}

	errUser := repository.UserRepository.FindUserByEmail(&user, forgotPassReq.Email).Error

	if errors.Is(errUser, gorm.ErrRecordNotFound) {
		return nil, &facades.ResponseError{
			StatusCode: facades.AppErrNotFound,
			Message:    translation.Lang(config.UserErrNotFound),
		}
	}

	errAct := repository.UserActionCodeRepository.FindUserActionCode(&userAct, user.ID.String(), forgotPassReq.Code).Error

	if errors.Is(errAct, gorm.ErrRecordNotFound) {
		return nil, &facades.ResponseError{
			StatusCode: facades.AppErrNotFound,
			Message:    translation.Lang(config.UserErrActivationNotFound),
		}
	}

	if userAct.IsUsed {
		return nil, &facades.ResponseError{
			StatusCode: facades.AppErrUnprocessable,
			Message:    translation.Lang(config.UserErrCodeAlreadyUsed),
		}
	}

	t := time.Now()

	if userAct.ExpiredAt.Before(t) {
		return nil, &facades.ResponseError{
			StatusCode: facades.AppErrExpire,
			Message:    translation.Lang(config.UserErrActivationExpired),
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

		repository.UserRepository.UpdatePassword(&userData)
		_, err := repository.RepositoryAggregate.UpdateActionCodeUsed(user.ID.String(), forgotPassReq.Code)
		if err != nil {
			return
		}
	}()

	return nil, nil
}
