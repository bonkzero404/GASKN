package contracts

import (
	"gorm.io/gorm"

	"gaskn/database/stores"
)

type UserInvitationRepository interface {
	FindUserInvitation(userInvitation *stores.UserInvitation, userId string, clientId string) *gorm.DB

	CreateUserInvitation(userInvitation *stores.UserInvitation) *gorm.DB

	UpdateUserInvitation(userInvitation *stores.UserInvitation) *gorm.DB
}
