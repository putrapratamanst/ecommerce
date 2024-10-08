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
	rabbitmq, err := messaging.NewRabbitMQ("amqp://user:password@rabbitmq:5672/")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect to RabbitMQ")
	}
	defer rabbitmq.Close()

	// Initialize services
	orderService := services.NewOrderService(orderRepo, rabbitmq)

	// Initialize controllers
	productServiceURL := os.Getenv("PRODUCT_SERVICE_URL")
	productServiceClient := services.NewProductServiceClient(productServiceURL)

	orderController := controllers.NewOrderController(orderService, productServiceClient)
	routes.SetupOrderRoutes(app, orderController)

	// Start the server
	port := os.Getenv("ORDER_SERVICE_PORT")
	app.Listen(":" + port) // Run the Fiber server
}
