package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/putrapratamanst/ecommerce/product-service/config"
	"github.com/putrapratamanst/ecommerce/product-service/routes"
)

func main() {
	// Initialize Fiber app
	app := fiber.New()

	// Load environment variables
	config.LoadEnv()

	// Initialize DB
	db := config.InitDB()

	// Initialize Redis
	redisClient := config.InitRedis()

	// Setup routes
	routes.SetupRoutes(app, db, redisClient)

	// Start the server
	port := os.Getenv("PRODUCT_SERVICE_PORT")
	app.Listen(":" + port)
}
