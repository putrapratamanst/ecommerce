package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/putrapratamanst/ecommerce/order-service/controllers"
	"github.com/putrapratamanst/ecommerce/order-service/middleware"
)

func SetupOrderRoutes(app *fiber.App, orderController *controllers.OrderController) {
	orderGroup := app.Group("/orders")
	{
		orderGroup.Post("/checkout", middleware.AuthMiddleware, orderController.CheckoutOrder)
		orderGroup.Post("/payment/confirm", middleware.AuthMiddleware, orderController.PaymentConfirm)
		orderGroup.Post("/cancel", middleware.AuthMiddleware, orderController.CancelOrder)
	}
}
