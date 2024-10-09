package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/putrapratamanst/ecommerce/warehouse-service/controllers"
	"github.com/putrapratamanst/ecommerce/warehouse-service/middleware"
)

func SetupWarehouseRoutes(app *fiber.App, warehouseController *controllers.WarehouseController) {

	app.Post("/warehouses", middleware.AuthMiddleware, warehouseController.CreateWarehouse)
	app.Post("/warehouses/:warehouseID/activate", warehouseController.ActivateWarehouse)
	app.Post("/warehouses/:warehouseID/shop/:shopID", middleware.AuthMiddleware, warehouseController.SetWarehouseShop)
	app.Post("/warehouses/transfer", middleware.AuthMiddleware, warehouseController.TransferStock)
	app.Put("/warehouses/:warehouseID/product/:productID/adjust", middleware.AuthMiddleware, warehouseController.AdjustStock)
}
