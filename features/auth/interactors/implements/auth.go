package implements

import (
	"errors"
	"github.com/bonkzero404/gaskn/config"
	"github.com/bonkzero404/gaskn/database/stores"
	responseDto "github.com/bonkzero404/gaskn/dto"
	"github.com/bonkzero404/gaskn/features/auth/dto"
	"github.com/bonkzero404/gaskn/features/auth/interactors"
	roleRepository "github.com/bonkzero404/gaskn/features/role/repositories"
	userInterface "github.com/bonkzero404/gaskn/features/user/repositories"
	"github.com/bonkzero404/gaskn/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

type Auth struct {
	UserRepository userInterface.UserRepository
	RoleRepository roleRepository.RoleRepository
}

func NewAuth(
	userRepository userInterface.UserRepository,
	roleRepository roleRepository.RoleRepository,
) interactors.UserAuth {
	return &Auth{
		UserRepository: userRepository,
		RoleRepository: roleRepository,
	}
}

func (service Auth) SetTokenResponse(c *fiber.Ctx, user *stores.User) (*dto.UserAuthResponse, error) {
	token, exp, errToken := utils.CreateToken(user.ID.String(), user.FullName)

	if errToken != nil {
		return nil, &responseDto.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    utils.Lang(c, config.AuthErrToken),
		}
	}

	// Set response message to succeed
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
func (service Auth) Authenticate(c *fiber.Ctx, auth *dto.UserAuthRequest) (*dto.UserAuthResponse, error) {
	var user stores.User

	// Get user by email
	errUser := service.UserRepository.FindUserByEmail(&user, auth.Email).Error

	// Check if the user is not found
	// then display an error message
	if errors.Is(errUser, gorm.ErrRecordNotFound) {
		return nil, &responseDto.ApiErrorResponse{
			StatusCode: fiber.StatusForbidden,
			Message:    utils.Lang(c, config.AuthErrGetProfile),
		}
	}

	// Check if a query operation error occurs
	if errUser != nil {
		return nil, &responseDto.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    utils.Lang(c, config.AuthErrRefreshToken),
		}
	}

	// Check if the user status is not active
	if !user.IsActive {
		return nil, &responseDto.ApiErrorResponse{
			StatusCode: fiber.StatusForbidden,
			Message:    utils.Lang(c, config.AuthErruserNotActive),
		}
	}

	// Match password hashes
	match := utils.CheckPasswordHash(auth.Password, user.Password)

	// Check if it doesn't match, show an error message
	if !match {
		return nil, &responseDto.ApiErrorResponse{
			StatusCode: fiber.StatusForbidden,
			Message:    utils.Lang(c, config.AuthErrInvalid),
		}
	}

	response, errResp := service.SetTokenResponse(c, &user)

	if errResp != nil {
		return nil, errResp
	}

	return response, nil
}

// GetProfile /*
func (service Auth) GetProfile(c *fiber.Ctx, id string) (*dto.UserAuthProfileResponse, error) {
	var user stores.User
	// var roleUser []stores.RoleUser

	// Get user from database
	errUser := service.UserRepository.FindUserById(&user, id).Error

	// Check if there is a query error
	if errUser != nil {
		return nil, &responseDto.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    utils.Lang(c, config.GlobalErrUnknown),
		}
	}

	// Set response message
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
func (service Auth) RefreshToken(c *fiber.Ctx, tokenUser *jwt.Token) (*dto.UserAuthResponse, error) {
	var user stores.User

	// Get data from token then convert to string
	beforeClaims := tokenUser.Claims.(jwt.MapClaims)
	id := beforeClaims["id"].(string)

	// Get user data
	errUser := service.UserRepository.FindUserById(&user, id).Error

	// Check if something went wrong with query
	if errUser != nil {
		return nil, &responseDto.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    utils.Lang(c, config.GlobalErrUnknown),
		}
	}

	response, errResp := service.SetTokenResponse(c, &user)

	if errResp != nil {
		return nil, errResp
	}

	return response, nil
}
