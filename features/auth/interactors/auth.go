package interactors

import (
	"github.com/bonkzero404/gaskn/database/stores"
	"github.com/bonkzero404/gaskn/features/auth/dto"

	"github.com/golang-jwt/jwt/v4"
)

type UserAuth interface {
	SetTokenResponse(user *stores.User) (*dto.UserAuthResponse, error)

	Authenticate(auth *dto.UserAuthRequest) (*dto.UserAuthResponse, error)

	GetProfile(id string) (*dto.UserAuthProfileResponse, error)

	RefreshToken(token *jwt.Token) (*dto.UserAuthResponse, error)
}
