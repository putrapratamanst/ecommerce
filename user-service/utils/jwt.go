package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/putrapratamanst/ecommerce/user-service/config"
	"github.com/putrapratamanst/ecommerce/user-service/models"
)

func GenerateToken(user *models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"role":  user.Role,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(config.JWT_SECRET)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
