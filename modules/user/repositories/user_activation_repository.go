package repositories

import (
	"go-starterkit-project/database/driver"
	"go-starterkit-project/database/stores"
	"go-starterkit-project/modules/user/contracts"

	"gorm.io/gorm"
)

type UserActivationRepository struct {
	DB *gorm.DB
}

func NewUserActivationRepository(db *gorm.DB) contracts.UserActivationRepositoryInterface {
	return &UserActivationRepository{
		DB: driver.DB,
	}
}

func (repository UserActivationRepository) FindUserActivationCode(
	userActivation *stores.UserActivation,
	userId string,
	code string,
) *gorm.DB {
	return repository.DB.First(&userActivation, "user_id = ? AND code = ?", userId, code)
}

func (repository UserActivationRepository) CreateUserActivation(userActivate *stores.UserActivation) *gorm.DB {
	return repository.DB.Create(&userActivate)
}

func (repository UserActivationRepository) UpdateActivationCodeUsed(userActivate *stores.UserActivation) *gorm.DB {
	return repository.DB.Save(&userActivate)
}
