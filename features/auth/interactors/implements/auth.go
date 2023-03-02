package implements

import (
	"errors"
	"github.com/bonkzero404/gaskn/app/facades"
	"github.com/bonkzero404/gaskn/app/translation"
	"github.com/bonkzero404/gaskn/app/utils"
	"github.com/bonkzero404/gaskn/config"
	"github.com/bonkzero404/gaskn/database/stores"
	"github.com/bonkzero404/gaskn/features/auth/dto"
	"github.com/bonkzero404/gaskn/features/auth/interactors"
	roleRepo "github.com/bonkzero404/gaskn/features/role/repositories"
	userRepo "github.com/bonkzero404/gaskn/features/user/repositories"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

type Auth struct {
	UserRepository userRepo.UserRepository
	RoleRepository roleRepo.RoleRepository
}

func NewAuth(
	userRepository userRepo.UserRepository,
	roleRepository roleRepo.RoleRepository,
) interactors.UserAuth {
	return &Auth{
		UserRepository: userRepository,
		RoleRepository: roleRepository,
	}
}

func (repository Auth) SetTokenResponse(user *stores.User) (*dto.UserAuthResponse, error) {
	token, exp, errToken := utils.CreateToken(user.ID.String(), user.FullName)

	if errToken != nil {
		return nil, &facades.ResponseError{
			StatusCode: facades.AppErrUnprocessable,
			Message:    translation.Lang(config.AuthErrToken),
		}
	}

	// Set facades message to succeed
	response := dto.UserAuthResponse{
		ID:       user.ID.String(),
		FullName: user.FullName,
		Email:    user.Email,
		Phone:    user.Phone,
		IsActive: user.IsActive,
		Token:    token,
		Exp:      exp,
	}

	return &response, nil
}

// Authenticate /*
func (repository Auth) Authenticate(auth *dto.UserAuthRequest) (*dto.UserAuthResponse, error) {
	var user stores.User

	// Get user by email
	errUser := repository.UserRepository.FindUserByEmail(&user, auth.Email).Error

	// Check if the user is not found
	// then display an error message
	if errors.Is(errUser, gorm.ErrRecordNotFound) {
		return nil, &facades.ResponseError{
			StatusCode: facades.AppErrNotFound,
			Message:    translation.Lang(config.AuthErrGetProfile),
		}
	}

	// Check if a query operation error occurs
	if errUser != nil {
		return nil, &facades.ResponseError{
			StatusCode: facades.AppErrUnprocessable,
			Message:    translation.Lang(config.AuthErrRefreshToken),
		}
	}

	// Check if the user status is not active
	if !user.IsActive {
		return nil, &facades.ResponseError{
			StatusCode: facades.AppErrForbidden,
			Message:    translation.Lang(config.AuthErruserNotActive),
		}
	}

	// Match password hashes
	match := utils.CheckPasswordHash(auth.Password, user.Password)

	// Check if it doesn't match, show an error message
	if !match {
		return nil, &facades.ResponseError{
			StatusCode: facades.AppErrForbidden,
			Message:    translation.Lang(config.AuthErrInvalid),
		}
	}

	response, errResp := repository.SetTokenResponse(&user)

	if errResp != nil {
		return nil, errResp
	}

	return response, nil
}

// GetProfile /*
func (repository Auth) GetProfile(id string) (*dto.UserAuthProfileResponse, error) {
	var user stores.User
	// var roleUser []stores.RoleUser

	// Get user from database
	errUser := repository.UserRepository.FindUserById(&user, id).Error

	// Check if there is a query error
	if errUser != nil {
		return nil, &facades.ResponseError{
			StatusCode: facades.AppErrUnprocessable,
			Message:    translation.Lang(config.GlobalErrUnknown),
		}
	}

	// Set facades message
	response := dto.UserAuthProfileResponse{
		ID:       user.ID.String(),
		FullName: user.FullName,
		Email:    user.Email,
		Phone:    user.Phone,
		IsActive: user.IsActive,
	}

	return &response, nil
}

// RefreshToken /*
func (repository Auth) RefreshToken(tokenUser *jwt.Token) (*dto.UserAuthResponse, error) {
	var user stores.User

	// Get data from token then convert to string
	beforeClaims := tokenUser.Claims.(jwt.MapClaims)
	id := beforeClaims["id"].(string)

	// Get user data
	errUser := repository.UserRepository.FindUserById(&user, id).Error

	// Check if something went wrong with query
	if errUser != nil {
		return nil, &facades.ResponseError{
			StatusCode: facades.AppErrUnprocessable,
			Message:    translation.Lang(config.GlobalErrUnknown),
		}
	}

	response, errResp := repository.SetTokenResponse(&user)

	if errResp != nil {
		return nil, errResp
	}

	return response, nil
}
