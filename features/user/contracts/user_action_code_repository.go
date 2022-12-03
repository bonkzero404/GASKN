package contracts

import (
	"gorm.io/gorm"

	"gaskn/database/stores"
)

type UserActionCodeRepository interface {
	FindUserActionCode(userActionCode *stores.UserActionCode, userId string, code string) *gorm.DB

	FindActionCode(userActivation *stores.UserActionCode, code string) *gorm.DB

	CreateUserActionCode(userActionCode *stores.UserActionCode) *gorm.DB

	UpdateActionCodeUsed(userActionCode *stores.UserActionCode) *gorm.DB
}
