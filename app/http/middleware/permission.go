package middleware

import (
	"errors"
	"github.com/bonkzero404/gaskn/app/http"
	"github.com/bonkzero404/gaskn/app/http/response"
	"github.com/bonkzero404/gaskn/app/translation"
	"github.com/bonkzero404/gaskn/config"
	"github.com/bonkzero404/gaskn/database/stores"
	"github.com/bonkzero404/gaskn/infrastructures"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

func Permission() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		enforcer := infrastructures.Enforcer

		var roleUser stores.RoleUser
		var roleUserClient stores.RoleUserClient
		// var client stores.Client
		var permit string

		token := c.Locals("user").(*jwt.Token)
		claims := token.Claims.(jwt.MapClaims)
		userId := claims["id"].(string)

		clientId := c.Params(config.Config("API_CLIENT_PARAM"))

		err := infrastructures.DB.Take(&roleUser, "user_id = ?", userId).Error

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.ApiForbidden(c, http.SetErrors{
				Message: translation.Lang(c, "middleware:err:unauthorized"),
				Cause:   "Forbidden access",
				Inputs:  nil,
			})
		} else {
			permit = "*"
		}

		errUserClient := infrastructures.DB.Take(&roleUserClient, "role_user_id = ? AND client_id = ?", roleUser.ID, clientId).Error

		if errors.Is(errUserClient, gorm.ErrRecordNotFound) {
			permit = "*"
		} else {
			permit = roleUserClient.ClientId.String()
		}

		if ok, _ := enforcer.Enforce(userId, permit, c.Path(), c.Method()); !ok {
			return response.ApiForbidden(c, http.SetErrors{
				Message: translation.Lang(c, "middleware:err:unauthorized"),
				Cause:   "Forbidden access",
				Inputs:  nil,
			})
		}

		return c.Next()
	}
}
