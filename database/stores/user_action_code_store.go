package stores

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ActCodeType string

// ACTIVATION_CODE /*
//
//goland:noinspection GoSnakeCaseUsage
const ACTIVATION_CODE ActCodeType = "ACTIVATION_CODE"

//goland:noinspection GoSnakeCaseUsage
const FORGOT_PASSWORD ActCodeType = "FORGOT_PASSWORD"

//goland:noinspection GoSnakeCaseUsage
const INVITATION_CODE ActCodeType = "INVITATION_CODE"

// UserActionCode /*
type UserActionCode struct {
	gorm.Model
	ID        uuid.UUID   `gorm:"type:char(36);primary_key"`
	UserId    uuid.UUID   `gorm:"type:char(36):index"`
	User      User        `gorm:"references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Code      string      `gorm:"type:char(32);index;not null"`
	ActType   ActCodeType `gorm:"type:char(30);index;not null"`
	ExpiredAt *time.Time
	IsUsed    bool
}

// BeforeCreate /*
func (*UserActionCode) BeforeCreate(tx *gorm.DB) error {
	t := time.Now()
	newT := t.Add(time.Hour * 2)

	tx.Statement.SetColumn("ID", uuid.New())
	tx.Statement.SetColumn("ExpiredAt", newT)
	return nil
}
