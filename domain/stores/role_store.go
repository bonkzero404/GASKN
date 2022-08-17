package stores

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

const (
	SA string = "sa"
	CL string = "cl"
)

/*
 *
 * Table model
 */
type Role struct {
	gorm.Model
	ID              uuid.UUID `gorm:"type:char(36);primary_key"`
	RoleName        string    `gorm:"type:varchar(100);index;not null"`
	RoleDescription string    `gorm:"type:text"`
	RoleType        string    `gorm:"type:char(2)"`
	IsActive        bool
}

/*
*
This function is a feature that gorm has for making hooks,
this hook function is used to generate uuid when the user
performs the create action
*/
func (*Role) BeforeCreate(tx *gorm.DB) error {
	tx.Statement.SetColumn("ID", uuid.New())
	return nil
}
