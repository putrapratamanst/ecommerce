package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/putrapratamanst/ecommerce/user-service/utils"
)

func (ctrl *AuthController)GetMe(c *fiber.Ctx) error {
    userID := c.Locals("userID")
	
    user, err := ctrl.AuthService.Detail(userID.(string))
	if err != nil {
		return utils.SendResponse(c, fiber.StatusNotFound, err.Error(), nil)
	}

	return utils.SendResponse(c, fiber.StatusCreated, "Successfully get detail user", fiber.Map{"user": user})
}