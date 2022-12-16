package interactors

import (
	"github.com/gofiber/fiber/v2"

	"github.com/bonkzero404/gaskn/database/stores"
	"github.com/bonkzero404/gaskn/features/user/dto"
)

type User interface {
	CreateUser(c *fiber.Ctx, user *dto.UserCreateRequest, isInternalRegister bool) (*dto.UserCreateResponse, error)

	UserActivation(c *fiber.Ctx, code string) (*dto.UserCreateResponse, error)

	CreateUserActivation(c *fiber.Ctx, email string, actType stores.ActCodeType) (any, error)

	UpdatePassword(c *fiber.Ctx, forgotPassReq *dto.UserForgotPassActRequest) (any, error)
}
