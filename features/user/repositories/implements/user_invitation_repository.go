package implements

import (
	"gaskn/features/user/repositories"
	"gorm.io/gorm"

	"gaskn/database/stores"
)

type UserInvitationRepository struct {
	DB *gorm.DB
}

func NewUserInvitationRepository(db *gorm.DB) repositories.UserInvitationRepository {
	return &UserInvitationRepository{
		DB: db,
	}
}

func (repository UserInvitationRepository) FindUserInvitation(userInvitation *stores.UserInvitation, userId string, clientId string) *gorm.DB {
	return repository.DB.Take(&userInvitation, "user_id = ? AND client_id = ?", userId, clientId)
}

func (repository UserInvitationRepository) FindInvitationByActId(userInvitation *stores.UserInvitation, actId string) *gorm.DB {
	return repository.DB.Take(&userInvitation, "user_action_code_id = ?", actId)
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
