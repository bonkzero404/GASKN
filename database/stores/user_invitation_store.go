package stores

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StatusInvitationType string

const (
	PENDING  StatusInvitationType = "pending"
	APPROVED StatusInvitationType = "approved"
	REJECTED StatusInvitationType = "rejected"
)

/*
*
Table model
*/
type UserInvitation struct {
	gorm.Model
	ID               uuid.UUID      `gorm:"type:char(36);primary_key"`
	UserId           uuid.UUID      `gorm:"type:char(36):index"`
	User             User           `gorm:"references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ClientId         uuid.UUID      `gorm:"type:char(36):index"`
	Client           Client         `gorm:"references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserActionCodeId uuid.UUID      `gorm:"type:char(36):index"`
	UserActionCode   UserActionCode `gorm:"references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	InvitedBy        string         `gorm:"type:varchar(100);not null"`
	UrlFrontendMatch string         `gorm:"type:text;not null"`
	Status           StatusInvitationType
}

/*
*
This function is a feature that gorm has for making hooks,
this hook function is used to generate uuid and add 2 hours
when the user performs the create action
*/
func (*UserInvitation) BeforeCreate(tx *gorm.DB) error {
	tx.Statement.SetColumn("ID", uuid.New())
	return nil
}
