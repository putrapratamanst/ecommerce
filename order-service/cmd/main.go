package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/putrapratamanst/ecommerce/order-service/config"
	"github.com/putrapratamanst/ecommerce/order-service/controllers"
	"github.com/putrapratamanst/ecommerce/order-service/messaging"
	"github.com/putrapratamanst/ecommerce/order-service/repositories"
	"github.com/putrapratamanst/ecommerce/order-service/routes"
	"github.com/putrapratamanst/ecommerce/order-service/services"
)

func main() {
	app := fiber.New()

	// Load environment variables
	config.LoadEnv()

	// Initialize DB
	db := config.InitDB()

	// Initialize repositories
	orderRepo := repositories.NewOrderRepository(db)

	// Initialize RabbitMQ
	rabbitMQ, err := messaging.NewRabbitMQ("amqp://user:password@rabbitmq:5672/")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect to RabbitMQ")
	}

	// Initialize services
	orderService := services.NewOrderService(orderRepo, rabbitMQ)

	// Initialize controllers
	orderController := controllers.NewOrderController(orderService)
	routes.SetupOrderRoutes(app, orderController)

	// Start the server
	port := os.Getenv("ORDER_SERVICE_PORT")
	app.Listen(":" + port) // Run the Fiber server
}
