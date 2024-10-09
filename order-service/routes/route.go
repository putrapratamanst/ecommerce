package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/putrapratamanst/ecommerce/order-service/controllers"
)

func SetupOrderRoutes(app *fiber.App, orderController *controllers.OrderController) {
	orderGroup := app.Group("/orders")
	{
		orderGroup.Post("/checkout", orderController.CheckoutOrder)
		orderGroup.Post("/payment/confirm", orderController.PaymentConfirm)
		orderGroup.Post("/cancel", orderController.CancelOrder)
	}
}
