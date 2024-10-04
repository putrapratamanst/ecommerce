package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/putrapratamanst/ecommerce/warehouse-service/config"
	"github.com/putrapratamanst/ecommerce/warehouse-service/controllers"
	"github.com/putrapratamanst/ecommerce/warehouse-service/repositories"
	"github.com/putrapratamanst/ecommerce/warehouse-service/routes"
	"github.com/putrapratamanst/ecommerce/warehouse-service/services"
)

func main() {
	// Initialize Fiber app
	app := fiber.New()

	// Load environment variables
	config.LoadEnv()

	// Initialize DB
	db := config.InitDB()

	// Setup routes
	warehouseRepo := repositories.NewWarehouseRepository(db)
	warehouseService := services.NewWarehouseService(warehouseRepo)
	warehouseController := controllers.NewWarehouseController(warehouseService)
	routes.SetupWarehouseRoutes(app, warehouseController)

	// Start the server
	port := os.Getenv("WAREHOUSE_SERVICE_PORT")
	app.Listen(":" + port)
}
