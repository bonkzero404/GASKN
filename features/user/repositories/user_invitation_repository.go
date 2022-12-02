package repositories

import (
	"gorm.io/gorm"

	"gaskn/database/driver"
	"gaskn/database/stores"
	"gaskn/features/user/contracts"
)

type UserInvitationRepository struct {
	DB *gorm.DB
}

func NewUserInvitationRepository(db *gorm.DB) contracts.UserInvitationRepository {
	return &UserInvitationRepository{
		DB: driver.DB,
	}
}

func (repository UserInvitationRepository) FindUserInvitation(userInvitation *stores.UserInvitation, userId string, clientId string) *gorm.DB {
	return repository.DB.First(&userInvitation, "user_id = ? AND client_id = ?", userId, clientId)
}

func (repository UserInvitationRepository) CreateUserInvitation(userInvitation *stores.UserInvitation) *gorm.DB {
	return repository.DB.Create(&userInvitation)
}

func (repository UserInvitationRepository) UpdateUserInvitation(userInvitation *stores.UserInvitation) *gorm.DB {
	return repository.DB.Save(&userInvitation)
}

func (repository UserInvitationRepository) CreateClientAssignment(clientAssign *stores.ClientAssignment) *gorm.DB {
	return repository.DB.Create(&clientAssign)
}
