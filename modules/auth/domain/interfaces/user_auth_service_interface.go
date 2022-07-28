package interfaces

import (
	"go-starterkit-project/modules/auth/domain/dto"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type UserAuthServiceInterface interface {
	Authenticate(c *fiber.Ctx, auth *dto.UserAuthRequest) (*dto.UserAuthResponse, error)

	GetProfile(c *fiber.Ctx, id string) (*dto.UserAuthProfileResponse, error)

	RefreshToken(c *fiber.Ctx, token *jwt.Token) (*dto.UserAuthResponse, error)
}
