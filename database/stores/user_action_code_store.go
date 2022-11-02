package stores

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ActCodeType string

/*
*
This constant is used to create a string enumeration type to
distinguish the type of user who is activating or requesting
forgotten passwords
*/
const (
	ACTIVATION_CODE ActCodeType = "ACTIVATION_CODE"
	FORGOT_PASSWORD ActCodeType = "FORGOT_PASSWORD"
	INVITATION_CODE ActCodeType = "INVITATION_CODE"
)

/*
*
Table model
*/
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

/*
*
This function is a feature that gorm has for making hooks,
this hook function is used to generate uuid and add 2 hours
when the user performs the create action
*/
func (*UserActionCode) BeforeCreate(tx *gorm.DB) error {
	t := time.Now()
	newT := t.Add(time.Hour * 2)

	tx.Statement.SetColumn("ID", uuid.New())
	tx.Statement.SetColumn("ExpiredAt", newT)
	return nil
}
