package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/putrapratamanst/ecommerce/shop-service/controllers"
	"github.com/putrapratamanst/ecommerce/shop-service/middleware"
)

func SetupShopRoutes(app *fiber.App, shopController *controllers.ShopController) {

	shopGroup := app.Group("/shops", middleware.AuthMiddleware)

	shopGroup.Post("/", shopController.CreateShop)
	shopGroup.Get("/:id", shopController.GetShop)
	shopGroup.Put("/:id", shopController.UpdateShop)
	shopGroup.Delete("/:id", shopController.DeleteShop)
	shopGroup.Get("/", shopController.GetAllShops)
}
