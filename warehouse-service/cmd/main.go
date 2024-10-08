package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/putrapratamanst/ecommerce/warehouse-service/config"
	"github.com/putrapratamanst/ecommerce/warehouse-service/controllers"
	"github.com/putrapratamanst/ecommerce/warehouse-service/messaging"
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

	// Initialize RabbitMQ
	rabbitMQ, err := messaging.NewRabbitMQ("amqp://user:password@rabbitmq:5672/")
	if err != nil {
		panic("failed to connect to RabbitMQ")
	}

	// Setup routes
	shopServiceURL := os.Getenv("SHOP_SERVICE_URL")

	warehouseRepo := repositories.NewWarehouseRepository(db)
	warehouseService := services.NewWarehouseService(warehouseRepo)
	shopServiceClient := services.NewShopServiceClient(shopServiceURL)
	warehouseController := controllers.NewWarehouseController(warehouseService, shopServiceClient)
	routes.SetupWarehouseRoutes(app, warehouseController)

	// Start listening for messages
	go warehouseService.StartListening(rabbitMQ)

	// Start the server
	port := os.Getenv("WAREHOUSE_SERVICE_PORT")
	app.Listen(":" + port)
}
