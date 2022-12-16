package implements

import (
	"github.com/bonkzero404/gaskn/database/stores"
	"github.com/bonkzero404/gaskn/features/user/repositories"
)

type RepositoryAggregate struct {
	UserRepository           repositories.UserRepository
	UserActionCodeRepository repositories.UserActionCodeRepository
}

func NewRepositoryAggregate(
	UserRepository repositories.UserRepository,
	UserActionCodeRepository repositories.UserActionCodeRepository,
) repositories.RepositoryAggregate {
	return &RepositoryAggregate{
		UserRepository:           UserRepository,
		UserActionCodeRepository: UserActionCodeRepository,
	}
}

func (repository RepositoryAggregate) CreateUser(user *stores.User, userActivate *stores.UserActionCode) (*stores.User, error) {
	if err := repository.UserRepository.CreateUser(user).Error; err != nil {
		return nil, err
	}

	userActivate.UserId = user.ID

	if err := repository.UserActionCodeRepository.CreateUserActionCode(userActivate).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (repository RepositoryAggregate) UpdateUserActivation(id string, stat bool) (*stores.User, error) {
	var user stores.User

	if err := repository.UserRepository.FindUserById(&user, id).Error; err != nil {
		return nil, err
	}

	user.IsActive = stat

	if err := repository.UserRepository.UpdateUserIsActive(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (repository RepositoryAggregate) UpdatePassword(id string, password string) (*stores.User, error) {
	var user stores.User

	if err := repository.UserRepository.FindUserById(&user, id).Error; err != nil {
		return nil, err
	}

	user.Password = password

	if err := repository.UserRepository.UpdatePassword(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (repository RepositoryAggregate) UpdateActionCodeUsed(userId string, code string) (*stores.UserActionCode, error) {
	var userAct stores.UserActionCode

	if err := repository.UserActionCodeRepository.FindUserActionCode(&userAct, userId, code).Error; err != nil {
		return nil, err
	}

	userAct.IsUsed = true

	if err := repository.UserActionCodeRepository.UpdateActionCodeUsed(&userAct).Error; err != nil {
		return nil, err
	}

	return &userAct, nil
}
