package contracts

import (
	"gaskn/modules/auth/dto"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type UserAuthService interface {
	Authenticate(c *fiber.Ctx, auth *dto.UserAuthRequest) (*dto.UserAuthResponse, error)

	GetProfile(c *fiber.Ctx, id string) (*dto.UserAuthProfileResponse, error)

	RefreshToken(c *fiber.Ctx, token *jwt.Token) (*dto.UserAuthResponse, error)
}
