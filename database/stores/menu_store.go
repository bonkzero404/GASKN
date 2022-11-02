package stores

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

const (
	MenuSA string = "sub"
	MenuCL string = "link"
)

/*
 *
 * Table model
 */
type Menu struct {
	gorm.Model
	ID              uuid.UUID `gorm:"type:char(36);primary_key"`
	ParentID        uuid.UUID `gorm:"type:char(36);index"`
	MenuName        string    `gorm:"type:varchar(100);index;not null"`
	MenuDescription string    `gorm:"type:text"`
	MenuUrl         string    `gorm:"type:text"`
	MenuType        string    `gorm:"type:char(4)"`
	IsActive        bool
}

/*
*
This function is a feature that gorm has for making hooks,
this hook function is used to generate uuid when the user
performs the create action
*/
func (*Menu) BeforeCreate(tx *gorm.DB) error {
	tx.Statement.SetColumn("ID", uuid.New())
	return nil
}
