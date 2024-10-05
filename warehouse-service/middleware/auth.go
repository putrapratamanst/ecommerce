package middleware

import (
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/putrapratamanst/ecommerce/warehouse-service/config"
	"github.com/putrapratamanst/ecommerce/warehouse-service/utils"
)

func AuthMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return utils.SendResponse(c, fiber.StatusUnauthorized, "Authorization header is required", nil)
	}

	tokenString := strings.Split(authHeader, "Bearer ")[1]
	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.JWT_SECRET), nil
	})
	if err != nil || !token.Valid {
		return utils.SendResponse(c, fiber.StatusUnauthorized, "Invalid token", nil)
	}

	role := claims["role"].(string)
	if role != "admin" {
		return utils.SendResponse(c, fiber.StatusUnauthorized, "User is not authorized", nil)
	}
	userIDFloat := claims["id"].(float64)
	userIDInt := int(userIDFloat)
	userID := strconv.Itoa(userIDInt)

	c.Locals("userID", userID)
	return c.Next()
}
