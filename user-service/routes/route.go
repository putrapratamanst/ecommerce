package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/putrapratamanst/ecommerce/user-service/controllers"
	"github.com/putrapratamanst/ecommerce/user-service/middleware"
)

func SetupAuthRoutes(app *fiber.App, authController *controllers.AuthController) {
	app.Post("/login", authController.Login)
	app.Post("/register", authController.Register)
	app.Get("/me", middleware.AuthMiddleware, authController.GetMe)
}
