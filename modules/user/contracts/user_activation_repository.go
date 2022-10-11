package contracts

import (
	"go-starterkit-project/database/stores"

	"gorm.io/gorm"
)

type UserActivationRepository interface {
	FindUserActivationCode(userActivation *stores.UserActivation, userId string, code string) *gorm.DB

	CreateUserActivation(userActivate *stores.UserActivation) *gorm.DB

	UpdateActivationCodeUsed(userActivation *stores.UserActivation) *gorm.DB
}
