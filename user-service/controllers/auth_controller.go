package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/putrapratamanst/ecommerce/user-service/models"
	"github.com/putrapratamanst/ecommerce/user-service/services"
)

type AuthController struct {
	AuthService services.AuthService
}

func NewAuthController(authService services.AuthService) *AuthController {
	return &AuthController{AuthService: authService}
}

func (ctrl *AuthController) Login(c *fiber.Ctx) error {
	var loginData struct {
		EmailOrPhone string `json:"email_or_phone"`
		Password     string `json:"password"`
	}

	if err := c.BodyParser(&loginData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	token, err := ctrl.AuthService.Login(loginData.EmailOrPhone, loginData.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"token": token})
}

func (ctrl *AuthController) Register(c *fiber.Ctx) error {
    var user models.User

    if err := c.BodyParser(&user); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
    }

    if err := ctrl.AuthService.Register(&user); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
    }

    return c.Status(fiber.StatusCreated).JSON(user)
}