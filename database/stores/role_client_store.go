package stores

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// RoleClient /*
type RoleClient struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:char(36);primary_key"`
	ClientId  uuid.UUID `gorm:"type:char(36):index"`
	Client    Client    `gorm:"references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	RoleId    uuid.UUID `gorm:"type:char(36):index"`
	Role      Role      `gorm:"references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CanDelete bool
	IsActive  bool
}

// BeforeCreate /*
func (*RoleClient) BeforeCreate(tx *gorm.DB) error {
	tx.Statement.SetColumn("ID", uuid.New())
	return nil
}
