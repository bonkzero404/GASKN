package services

import (
	"errors"
	"go-starterkit-project/database/stores"
	respModel "go-starterkit-project/dto"
	"go-starterkit-project/modules/auth/contracts"
	"go-starterkit-project/modules/auth/dto"
	roleRepository "go-starterkit-project/modules/role/contracts"
	userInterface "go-starterkit-project/modules/user/contracts"
	"go-starterkit-project/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

type AuthService struct {
	UserRepository userInterface.UserRepository
	RoleRepository roleRepository.RoleRepository
}

func NewAuthService(
	userRepository userInterface.UserRepository,
	roleRepository roleRepository.RoleRepository,
) contracts.UserAuthService {
	return &AuthService{
		UserRepository: userRepository,
		RoleRepository: roleRepository,
	}
}

/*
*
This function is used to handle authentication
*/
func (service AuthService) Authenticate(c *fiber.Ctx, auth *dto.UserAuthRequest) (*dto.UserAuthResponse, error) {
	var user stores.User

	// Get user by email
	errUser := service.UserRepository.FindUserByEmail(&user, auth.Email).Error

	// Check if the user is not found
	// then displayan error message
	if errors.Is(errUser, gorm.ErrRecordNotFound) {
		return &dto.UserAuthResponse{}, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusForbidden,
			Message:    utils.Lang(c, "auth:err:invalid-auth"),
		}
	}

	// Check if a query operation error occurs
	if errUser != nil {
		return &dto.UserAuthResponse{}, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    utils.Lang(c, "global:err:failed-unknown"),
		}
	}

	// Check if the user status is not active
	if !user.IsActive {
		return &dto.UserAuthResponse{}, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusForbidden,
			Message:    utils.Lang(c, "auth:err:user-not-active"),
		}
	}

	// Match password hashes
	match := utils.CheckPasswordHash(auth.Password, user.Password)

	// Check if it doesn't match, show an error message
	if !match {
		return &dto.UserAuthResponse{}, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusForbidden,
			Message:    utils.Lang(c, "auth:err:invalid-auth"),
		}
	}

	token, exp, errToken := utils.CreateToken(user.ID.String())

	if errToken != nil {
		return &dto.UserAuthResponse{}, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    utils.Lang(c, "auth:err:err-token"),
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

/*
*
This function is used to authorize users and display logged in user data
*/
func (service AuthService) GetProfile(c *fiber.Ctx, id string) (*dto.UserAuthProfileResponse, error) {
	var user stores.User
	// var roleUser []stores.RoleUser

	// Get user from database
	errUser := service.UserRepository.FindUserById(&user, id).Error

	// Check if there is a query error
	if errUser != nil {
		return &dto.UserAuthProfileResponse{}, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    utils.Lang(c, "global:err:failed-unknown"),
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

/*
*
This function is used to refresh token
*/
func (service AuthService) RefreshToken(c *fiber.Ctx, tokenUser *jwt.Token) (*dto.UserAuthResponse, error) {
	var user stores.User

	// Get data from token then convert to string
	beforeClaims := tokenUser.Claims.(jwt.MapClaims)
	id := beforeClaims["id"].(string)

	// Get user data
	errUser := service.UserRepository.FindUserById(&user, id).Error

	// Check if something went wrong with query
	if errUser != nil {
		return &dto.UserAuthResponse{}, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    utils.Lang(c, "global:err:failed-unknown"),
		}
	}

	token, exp, errToken := utils.CreateToken(user.ID.String())
	if errToken != nil {
		return &dto.UserAuthResponse{}, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    utils.Lang(c, "auth:err:err-token"),
		}
	}

	// Set response message
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
