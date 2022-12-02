package stores

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

const (
	SA string = "sa"
	CL string = "cl"
)

// Role /*
type Role struct {
	gorm.Model
	ID              uuid.UUID `gorm:"type:char(36);primary_key"`
	RoleName        string    `gorm:"type:varchar(100);index;not null"`
	RoleDescription string    `gorm:"type:text"`
	RoleType        string    `gorm:"type:char(2)"`
	IsActive        bool
}

// BeforeCreate /*
func (*Role) BeforeCreate(tx *gorm.DB) error {
	tx.Statement.SetColumn("ID", uuid.New())
	return nil
}
