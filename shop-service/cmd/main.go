package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/putrapratamanst/ecommerce/shop-service/config"
	"github.com/putrapratamanst/ecommerce/shop-service/controllers"
	"github.com/putrapratamanst/ecommerce/shop-service/repositories"
	"github.com/putrapratamanst/ecommerce/shop-service/routes"
	"github.com/putrapratamanst/ecommerce/shop-service/services"
)

func main() {
	// Initialize Fiber app
	app := fiber.New()

	// Load environment variables
	config.LoadEnv()

	// Initialize DB
	db := config.InitDB()

	// Setup routes
	shopRepo := repositories.NewShopRepository(db)
	shopService := services.NewShopService(shopRepo)
    shopController := controllers.NewShopController(shopService)
    routes.SetupShopRoutes(app, shopController)

	// Start the server
	port := os.Getenv("SHOP_SERVICE_PORT")
	app.Listen(":" + port)
}
