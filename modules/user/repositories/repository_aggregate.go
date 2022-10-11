package repositories

import (
	"go-starterkit-project/database/stores"
	"go-starterkit-project/modules/user/contracts"
)

type RepositoryAggregate struct {
	UserRepository           contracts.UserRepositoryInterface
	UserActivationRepository contracts.UserActivationRepositoryInterface
}

func NewRepositoryAggregate(
	userRepository contracts.UserRepositoryInterface,
	userActivationRepository contracts.UserActivationRepositoryInterface,
) contracts.RepositoryAggregateInterface {
	return &RepositoryAggregate{
		UserRepository:           userRepository,
		UserActivationRepository: userActivationRepository,
	}
}

func (repository RepositoryAggregate) CreateUser(user *stores.User, userActivate *stores.UserActivation) (*stores.User, error) {
	if err := repository.UserRepository.CreateUser(user).Error; err != nil {
		return &stores.User{}, err
	}

	userActivate.UserId = user.ID

	if err := repository.UserActivationRepository.CreateUserActivation(userActivate).Error; err != nil {
		return &stores.User{}, err
	}

	return user, nil
}

func (repository RepositoryAggregate) UpdateUserActivation(id string, stat bool) (*stores.User, error) {
	var user stores.User

	if err := repository.UserRepository.FindUserById(&user, id).Error; err != nil {
		return &stores.User{}, err
	}

	user.IsActive = stat

	if err := repository.UserRepository.UpdateUserIsActive(&user).Error; err != nil {
		return &stores.User{}, err
	}

	return &user, nil
}

func (repository RepositoryAggregate) UpdatePassword(id string, password string) (*stores.User, error) {
	var user stores.User

	if err := repository.UserRepository.FindUserById(&user, id).Error; err != nil {
		return &stores.User{}, err
	}

	user.Password = password

	if err := repository.UserRepository.UpdatePassword(&user).Error; err != nil {
		return &stores.User{}, err
	}

	return &user, nil
}

func (repository RepositoryAggregate) UpdateActivationCodeUsed(userId string, code string) (*stores.UserActivation, error) {
	var userAct stores.UserActivation

	if err := repository.UserActivationRepository.FindUserActivationCode(&userAct, userId, code).Error; err != nil {
		return &stores.UserActivation{}, err
	}

	userAct.IsUsed = true

	if err := repository.UserActivationRepository.UpdateActivationCodeUsed(&userAct).Error; err != nil {
		return &stores.UserActivation{}, err
	}

	return &userAct, nil
}
