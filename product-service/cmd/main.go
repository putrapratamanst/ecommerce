package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/putrapratamanst/ecommerce/product-service/config"
	"github.com/putrapratamanst/ecommerce/product-service/routes"
	"github.com/putrapratamanst/ecommerce/product-service/services"
	"github.com/putrapratamanst/ecommerce/product-service/repositories"
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

	// Setup routes
	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo,redisClient)
    productController := controllers.NewProductController(productService)
	routes.SetupRoutes(app,productController)

	// Start the server
	port := os.Getenv("PRODUCT_SERVICE_PORT")
	app.Listen(":" + port)
}
