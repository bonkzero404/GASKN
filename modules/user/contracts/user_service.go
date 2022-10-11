package contracts

import (
	"gaskn/database/stores"
	"gaskn/modules/user/dto"

	"github.com/gofiber/fiber/v2"
)

type UserService interface {
	CreateUser(c *fiber.Ctx, user *dto.UserCreateRequest) (*dto.UserCreateResponse, error)

	UserActivation(c *fiber.Ctx, email string, code string) (*dto.UserCreateResponse, error)

	CreateUserActivation(c *fiber.Ctx, email string, actType stores.ActivationType) (map[string]interface{}, error)

	UpdatePassword(c *fiber.Ctx, forgotPassReq *dto.UserForgotPassActRequest) (map[string]interface{}, error)
}
