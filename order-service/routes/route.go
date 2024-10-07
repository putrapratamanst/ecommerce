package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/putrapratamanst/ecommerce/order-service/controllers"
)

func SetupOrderRoutes(app *fiber.App, orderController *controllers.OrderController) {
	orderGroup := app.Group("/orders")
	{
		orderGroup.Post("/create", orderController.PlaceOrder)
		orderGroup.Get("/:id", orderController.GetOrder)
		orderGroup.Get("/", orderController.GetOrders)
	}
}
