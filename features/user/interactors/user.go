package interactors

import (
	"github.com/bonkzero404/gaskn/database/stores"
	"github.com/bonkzero404/gaskn/features/user/dto"
)

type User interface {
	CreateUser(clientId string, user *dto.UserCreateRequest, isInternalRegister bool) (*dto.UserCreateResponse, error)

	UserActivation(code string) (*dto.UserCreateResponse, error)

	CreateUserAction(email string, actType stores.ActCodeType) (any, error)

	UpdatePassword(forgotPassReq *dto.UserForgotPassActRequest) (any, error)
}
