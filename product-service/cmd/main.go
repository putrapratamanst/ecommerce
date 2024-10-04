package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/putrapratamanst/ecommerce/product-service/config"
	"github.com/putrapratamanst/ecommerce/product-service/controllers"
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

	// Initialize Product controller
	productController := controllers.NewProductController(db, redisClient)

	// Setup routes
	app.Get("/products", productController.GetProducts)

	// Start the server
	port := os.Getenv("PRODUCT_SERVICE_PORT")
	app.Listen(":" + port)
}
