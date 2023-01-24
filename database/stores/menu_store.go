package stores

import (
	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type MenuType string

type LangAttribute struct {
	En string `json:"en"`
	Id string `json:"id"`
}

//goland:noinspection GoUnusedConst,GoUnusedConst
const (
	MenuBO MenuType = "BO" // BackOffice Type
	MenuCL MenuType = "CL" // Client Type
)

// Menu /*
type Menu struct {
	gorm.Model
	ID              uuid.UUID `gorm:"type:char(36);primary_key"`
	ParentID        uuid.UUID `gorm:"type:char(36);index"`
	MenuName        datatypes.JSONType[LangAttribute]
	MenuDescription datatypes.JSONType[LangAttribute]
	MenuUrl         string   `gorm:"type:text"`
	MenuIcon        string   `gorm:"menu_icon"`
	MenuType        MenuType `gorm:"type:char(2);index"`
	Sort            int
	IsActive        bool
}

// BeforeCreate /*
func (*Menu) BeforeCreate(tx *gorm.DB) error {
	tx.Statement.SetColumn("ID", uuid.New())
	return nil
}
