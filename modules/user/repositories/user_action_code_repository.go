package repositories

import (
	"gorm.io/gorm"

	"gaskn/database/driver"
	"gaskn/database/stores"
	"gaskn/modules/user/contracts"
)

type UserActionCodeRepository struct {
	DB *gorm.DB
}

func NewUserActionCodeRepository(db *gorm.DB) contracts.UserActionCodeRepository {
	return &UserActionCodeRepository{
		DB: driver.DB,
	}
}

func (repository UserActionCodeRepository) FindUserActionCode(
	userActivation *stores.UserActionCode,
	userId string,
	code string,
) *gorm.DB {
	return repository.DB.First(&userActivation, "user_id = ? AND code = ?", userId, code)
}

func (repository UserActionCodeRepository) CreateUserActionCode(userActivate *stores.UserActionCode) *gorm.DB {
	return repository.DB.Create(&userActivate)
}

func (repository UserActionCodeRepository) UpdateActionCodeUsed(userActivate *stores.UserActionCode) *gorm.DB {
	return repository.DB.Save(&userActivate)
}
