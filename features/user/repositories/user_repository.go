package repositories

import (
	"github.com/bonkzero404/gaskn/database/stores"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *stores.User) *gorm.DB

	UpdateUserIsActive(user *stores.User) *gorm.DB

	FindUserByEmail(user *stores.User, email string) *gorm.DB

	FindUserById(user *stores.User, id string) *gorm.DB

	UpdatePassword(user *stores.User) *gorm.DB
}
