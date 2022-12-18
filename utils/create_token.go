package utils

import (
	"github.com/bonkzero404/gaskn/config"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func CreateToken(id string, name string) (string, int64, error) {
	// Recreate token
	exp := time.Now().Add(time.Hour * 72).Unix()
	claims := jwt.MapClaims{
		"id":   id,
		"name": name,
		"exp":  exp,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, errToken := token.SignedString([]byte(config.Config("JWT_SECRET")))
	if errToken != nil {
		return "", 0, errToken
	}

	return t, exp, nil
}
