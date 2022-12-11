package implements

import (
	"gaskn/database/stores"
	"gaskn/features/user/repositories"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) repositories.UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (repository UserRepository) CreateUser(user *stores.User) *gorm.DB {
	return repository.DB.Create(&user)
}

func (repository UserRepository) UpdateUserIsActive(user *stores.User) *gorm.DB {
	return repository.DB.Save(&user)
}

func (repository UserRepository) FindUserByEmail(user *stores.User, email string) *gorm.DB {
	return repository.DB.Take(&user, "email = ?", email)
}

func (repository UserRepository) FindUserById(user *stores.User, id string) *gorm.DB {
	return repository.DB.Take(&user, "id = ?", id)
}

func (repository UserRepository) UpdatePassword(user *stores.User) *gorm.DB {
	return repository.DB.Save(&user)
}
