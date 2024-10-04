package controllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/putrapratamanst/ecommerce/user-service/models"
	"github.com/putrapratamanst/ecommerce/user-service/services"
	"github.com/putrapratamanst/ecommerce/user-service/utils"
)

type AuthController struct {
	AuthService services.AuthService
	validate    *validator.Validate
}

func NewAuthController(authService services.AuthService) *AuthController {
	return &AuthController{
		AuthService: authService,
		validate:    validator.New(),
	}
}

func (ctrl *AuthController) Login(c *fiber.Ctx) error {
	var loginData struct {
		EmailOrPhone string `json:"email_or_phone" validate:"required"`
		Password     string `json:"password" validate:"required"`
	}

	if err := c.BodyParser(&loginData); err != nil {
		return utils.SendResponse(c, fiber.StatusBadRequest, "Invalid input", nil)
	}

	if err := ctrl.validate.Struct(loginData); err != nil {
		return utils.SendResponse(c, fiber.StatusBadRequest, "Invalid input validation", nil)
	}

	token, err := ctrl.AuthService.Login(loginData.EmailOrPhone, loginData.Password)
	if err != nil {
		return utils.SendResponse(c, fiber.StatusUnauthorized, err.Error(), nil)
	}

	return utils.SendResponse(c, fiber.StatusOK, "Login successful", fiber.Map{"token": token})
}

func (ctrl *AuthController) Register(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return utils.SendResponse(c, fiber.StatusBadRequest, "Invalid input", nil)
	}

	if err := ctrl.validate.Struct(user); err != nil {
		return utils.SendResponse(c, fiber.StatusBadRequest, "Invalid input validation: "+err.Error(), nil)
	}

	if err := ctrl.AuthService.Register(&user); err != nil {
		return utils.SendResponse(c, fiber.StatusBadRequest, err.Error(), nil)
	}

	tokenString, err := utils.GenerateToken(&user)
	if err != nil {
		return utils.SendResponse(c, fiber.StatusBadRequest, err.Error(), nil)
	}

	return utils.SendResponse(c, fiber.StatusCreated, "User registered successfully", fiber.Map{"user": user, "acccessToken": tokenString})
}
