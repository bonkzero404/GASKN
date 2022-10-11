package contracts

import (
	"go-starterkit-project/database/stores"

	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	CreateUser(user *stores.User) *gorm.DB

	UpdateUserIsActive(user *stores.User) *gorm.DB

	FindUserByEmail(user *stores.User, email string) *gorm.DB

	FindUserById(user *stores.User, id string) *gorm.DB

	UpdatePassword(user *stores.User) *gorm.DB
}
