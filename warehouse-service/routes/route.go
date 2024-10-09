package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/putrapratamanst/ecommerce/warehouse-service/controllers"
	"github.com/putrapratamanst/ecommerce/warehouse-service/middleware"
)

func SetupWarehouseRoutes(app *fiber.App, warehouseController *controllers.WarehouseController) {

	app.Post("/warehouses", middleware.AuthMiddleware, warehouseController.CreateWarehouse)
	app.Post("/warehouses/:warehouseID/shop/:shopID", middleware.AuthMiddleware, warehouseController.SetWarehouseShop)
    app.Post("/warehouses/transfer", warehouseController.TransferStock)
    app.Put("/warehouse/:warehouseID/product/:productID/adjust", warehouseController.AdjustStock)
    app.Post("/warehouse/transfer", warehouseController.TransferStock)
    app.Post("/warehouse/:warehouseID/activate", warehouseController.ActivateWarehouse)
}