package stores

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// RoleUser /*
type RoleUser struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:char(36);primary_key"`
	UserId   uuid.UUID `gorm:"type:char(36):index"`
	RoleId   uuid.UUID `gorm:"type:char(36):index"`
	User     User      `gorm:"references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Role     Role      `gorm:"references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	IsActive bool
}

// BeforeCreate /*
func (*RoleUser) BeforeCreate(tx *gorm.DB) error {
	tx.Statement.SetColumn("ID", uuid.New())
	return nil
}
