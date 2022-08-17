package stores

import (
	"database/sql/driver"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type roleType string

const (
	SA roleType = "sa"
	CL roleType = "cl"
)

func (rt *roleType) Scan(value interface{}) error {
	*rt = roleType(value.([]byte))
	return nil
}

func (rt roleType) Value() (driver.Value, error) {
	return string(rt), nil
}

/*
 *
 * Table model
 */
type Role struct {
	gorm.Model
	ID              uuid.UUID `gorm:"type:char(36);primary_key"`
	RoleName        string    `gorm:"type:varchar(100);index;not null"`
	RoleDescription string    `gorm:"type:text"`
	RoleType        roleType  `sql:"role_type"`
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
