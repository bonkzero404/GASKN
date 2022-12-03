package stores

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// RoleUserClient /*
type RoleUserClient struct {
	gorm.Model
	ID         uuid.UUID `gorm:"type:char(36);primary_key"`
	ClientId   uuid.UUID `gorm:"type:char(36):index"`
	RoleUserId uuid.UUID `gorm:"type:char(36):index"`
	Client     Client    `gorm:"references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	RoleUser   RoleUser  `gorm:"references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	IsActive   bool
}

// BeforeCreate /*
func (*RoleUserClient) BeforeCreate(tx *gorm.DB) error {
	tx.Statement.SetColumn("ID", uuid.New())
	return nil
}
