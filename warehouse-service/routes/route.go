package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/putrapratamanst/ecommerce/warehouse-service/controllers"
	"github.com/putrapratamanst/ecommerce/warehouse-service/middleware"
)

func SetupWarehouseRoutes(app *fiber.App, warehouseController *controllers.WarehouseController) {

	app.Post("/warehouse", middleware.AuthMiddleware, warehouseController.CreateWarehouse)
	app.Post("/warehouse/:warehouseID/shop/:shopID", middleware.AuthMiddleware, warehouseController.SetWarehouseShop)
}