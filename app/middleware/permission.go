package middleware

import (
	"errors"
	"go-starterkit-project/config"
	"go-starterkit-project/database/driver"
	"go-starterkit-project/domain/dto"
	"go-starterkit-project/domain/stores"
	"go-starterkit-project/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

func Permission() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		enforcer := driver.Enforcer

		var roleUser stores.RoleUser
		// var client stores.Client
		var permit string

		token := c.Locals("user").(*jwt.Token)
		claims := token.Claims.(jwt.MapClaims)
		userId := claims["id"].(string)

		clientId := c.Params(config.Config("API_CLIENT_PARAM"))

		err := driver.DB.Preload("Client").Take(&roleUser, "user_id = ? AND client_id = ?", userId, clientId).Error

		if errors.Is(err, gorm.ErrRecordNotFound) {
			permit = "*"
		} else {
			permit = roleUser.ClientId.String()
		}

		if ok, _ := enforcer.Enforce(userId, permit, c.Path(), c.Method()); !ok {
			return utils.ApiForbidden(c, dto.Errors{
				Message: utils.Lang(c, "middleware:err:unauthorized"),
				Cause:   "Forbidden access",
				Inputs:  nil,
			})
		}

		return c.Next()
	}
}
