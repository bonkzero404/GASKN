package implements

import (
	"gaskn/features/user/repositories"
	"gorm.io/gorm"

	"gaskn/database/stores"
)

type UserActionCodeRepository struct {
	DB *gorm.DB
}

func NewUserActionCodeRepository(db *gorm.DB) repositories.UserActionCodeRepository {
	return &UserActionCodeRepository{
		DB: db,
	}
}

func (repository UserActionCodeRepository) FindUserActionCode(
	userActivation *stores.UserActionCode,
	userId string,
	code string,
) *gorm.DB {
	return repository.DB.First(&userActivation, "user_id = ? AND code = ?", userId, code)
}

func (repository UserActionCodeRepository) FindActionCode(userActivation *stores.UserActionCode, code string) *gorm.DB {
	return repository.DB.First(&userActivation, "code = ?", code)
}

func (repository UserActionCodeRepository) FindExistsActionCode(userActionCode *stores.UserActionCode, userId string, actType stores.ActCodeType) *gorm.DB {
	return repository.DB.Take(&userActionCode, "user_id = ? and act_type = ?", userId, actType)
}

func (repository UserActionCodeRepository) CreateUserActionCode(userActivate *stores.UserActionCode) *gorm.DB {
	return repository.DB.Create(&userActivate)
}

func (repository UserActionCodeRepository) UpdateActionCodeUsed(userActivate *stores.UserActionCode) *gorm.DB {
	return repository.DB.Save(&userActivate)
}
