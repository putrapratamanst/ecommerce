package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/putrapratamanst/ecommerce/product-service/controllers"
)

func SetupRoutes(app *fiber.App, productController *controllers.ProductController) {

	// example route: http://127.0.0.1:3000/products?page=2&limit=2
	app.Get("/products", productController.GetProducts)
	app.Get("/products/:id", productController.GetProductByID)

}
