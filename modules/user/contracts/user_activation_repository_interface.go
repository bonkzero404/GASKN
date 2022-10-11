package contracts

import (
	"go-starterkit-project/database/stores"

	"gorm.io/gorm"
)

type UserActivationRepositoryInterface interface {
	FindUserActivationCode(userActivation *stores.UserActivation, userId string, code string) *gorm.DB

	CreateUserActivation(userActivate *stores.UserActivation) *gorm.DB

	UpdateActivationCodeUsed(userActivation *stores.UserActivation) *gorm.DB
}
