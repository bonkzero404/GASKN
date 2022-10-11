package stores

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

/**
Table model
*/
type User struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:char(36);primary_key"`
	FullName string    `gorm:"type:varchar(100);index;not null"`
	Email    string    `gorm:"type:varchar(30);uniqueIndex;not null"`
	Phone    string    `gorm:"type:varchar(20);uniqueIndex;not null"`
	Password string    `gorm:"type:varchar(60);not null"`
	IsActive bool
}

/**
This function is a feature that gorm has for making hooks,
this hook function is used to generate uuid when the user
performs the create action
*/
func (*User) BeforeCreate(tx *gorm.DB) error {
	tx.Statement.SetColumn("ID", uuid.New())
	return nil
}
