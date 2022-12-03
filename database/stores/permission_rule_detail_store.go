package stores

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// PermissionRuleDetail /*
type PermissionRuleDetail struct {
	gorm.Model
	ID               uuid.UUID      `gorm:"type:char(36);primary_key"`
	PermissionRuleId uint           `gorm:"index"`
	PermissionRule   PermissionRule `gorm:"references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserName         string         `gorm:"type:varchar(255);index"`
	RoleName         string         `gorm:"type:varchar(255);index"`
	ClientName       string         `gorm:"type:varchar(255);index"`
}

// BeforeCreate /*
func (*PermissionRuleDetail) BeforeCreate(tx *gorm.DB) error {
	tx.Statement.SetColumn("ID", uuid.New())
	return nil
}
