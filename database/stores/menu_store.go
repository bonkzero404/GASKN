package stores

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MenuType string

const (
	MenuBO MenuType = "BO" // BackOffice Type
	MenuCL MenuType = "CL" // Client Type
)

// Menu /*
type Menu struct {
	gorm.Model
	ID              uuid.UUID `gorm:"type:char(36);primary_key"`
	ParentID        uuid.UUID `gorm:"type:char(36);index"`
	MenuName        string    `gorm:"type:varchar(100);index;not null"`
	MenuDescription string    `gorm:"type:text"`
	MenuUrl         string    `gorm:"type:text"`
	MenuType        MenuType  `gorm:"type:char(2);index"`
	IsActive        bool
}

// BeforeCreate /*
func (*Menu) BeforeCreate(tx *gorm.DB) error {
	tx.Statement.SetColumn("ID", uuid.New())
	return nil
}
