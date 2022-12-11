package repositories

import (
	"gorm.io/gorm"

	"gaskn/database/stores"
)

type UserInvitationRepository interface {
	FindUserInvitation(userInvitation *stores.UserInvitation, userId string, clientId string) *gorm.DB

	FindInvitationByActId(userInvitation *stores.UserInvitation, actId string) *gorm.DB

	CreateUserInvitation(userInvitation *stores.UserInvitation) *gorm.DB

	UpdateUserInvitation(userInvitation *stores.UserInvitation) *gorm.DB

	CreateClientAssignment(clientAssign *stores.ClientAssignment) *gorm.DB
}
