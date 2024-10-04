package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/putrapratamanst/ecommerce/user-service/config"
	"github.com/putrapratamanst/ecommerce/user-service/controllers"
	"github.com/putrapratamanst/ecommerce/user-service/repositories"
	"github.com/putrapratamanst/ecommerce/user-service/routes"
	"github.com/putrapratamanst/ecommerce/user-service/services"
)

func main() {
	// Initialize Fiber app
	app := fiber.New()

	// Load environment variables
	config.LoadEnv()

	// Initialize DB
	db := config.InitDB()

	// Setup routes
	userRepo := repositories.NewUserRepository(db)
	authService := services.NewAuthService(userRepo)
    authController := controllers.NewAuthController(authService)
    routes.SetupAuthRoutes(app, authController)

	// Start the server
	port := os.Getenv("USER_SERVICE_PORT")
	app.Listen(":" + port)
}
