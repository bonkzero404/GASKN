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

// UserInvitation /*
type UserInvitation struct {
	gorm.Model
	ID               uuid.UUID      `gorm:"type:char(36);primary_key"`
	UserId           uuid.UUID      `gorm:"type:char(36):index"`
	User             User           `gorm:"references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ClientId         uuid.UUID      `gorm:"type:char(36):index"`
	Client           Client         `gorm:"references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserActionCodeId uuid.UUID      `gorm:"type:char(36):index"`
	UserActionCode   UserActionCode `gorm:"references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	RoleClientId     uuid.UUID      `gorm:"type:char(36):index"`
	RoleClient       RoleClient     `gorm:"references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	InvitedBy        string         `gorm:"type:varchar(100);not null"`
	Role             string         `gorm:"type:varchar(100);not null"`
	UrlFrontendMatch string         `gorm:"type:text;not null"`
	Status           StatusInvitationType
}

// BeforeCreate /*
func (*UserInvitation) BeforeCreate(tx *gorm.DB) error {
	tx.Statement.SetColumn("ID", uuid.New())
	return nil
}
