package middleware

import (
	"errors"
	"fmt"
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
		var roleUserClient stores.RoleUserClient
		// var client stores.Client
		var permit string

		token := c.Locals("user").(*jwt.Token)
		claims := token.Claims.(jwt.MapClaims)
		userId := claims["id"].(string)

		clientId := c.Params(config.Config("API_CLIENT_PARAM"))

		err := driver.DB.Take(&roleUser, "user_id = ?", userId).Error

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.ApiForbidden(c, dto.Errors{
				Message: utils.Lang(c, "middleware:err:unauthorized"),
				Cause:   "Forbidden access",
				Inputs:  nil,
			})
		} else {
			permit = "*"
		}

		errUserClient := driver.DB.Take(&roleUserClient, "role_user_id = ? AND client_id = ?", roleUser.ID, clientId).Error

		if errors.Is(errUserClient, gorm.ErrRecordNotFound) {
			permit = "*"
		} else {
			permit = roleUserClient.ClientId.String()
		}

		fmt.Print(userId, permit, c.Path(), c.Method())

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
