package repositories

import (
	"gorm.io/gorm"

	"github.com/bonkzero404/gaskn/database/stores"
)

type UserActionCodeRepository interface {
	FindUserActionCode(userActionCode *stores.UserActionCode, userId string, code string) *gorm.DB

	FindActionCode(userActivation *stores.UserActionCode, code string) *gorm.DB

	FindExistsActionCode(userActionCode *stores.UserActionCode, userId string, actType stores.ActCodeType) *gorm.DB

	CreateUserActionCode(userActionCode *stores.UserActionCode) *gorm.DB

	UpdateActionCodeUsed(userActionCode *stores.UserActionCode) *gorm.DB
}
