package cmd

import (
	"github.com/gofiber/fiber/v2"
	"github.com/putrapratamanst/ecommerce/product-service/config"
)

func main() {
	// Initialize Fiber app
	app := fiber.New()

	// Initialize DB
	db := config.InitDB()

	// Initialize Product controller
	productController := controllers.NewProductController(db)

	// Setup routes
	app.Get("/products", productController.GetProducts)

	// Start the server
	app.Listen(":3000")
}
