package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/putrapratamanst/ecommerce/warehouse-service/controllers"
	"github.com/putrapratamanst/ecommerce/warehouse-service/middleware"
)

func SetupWarehouseRoutes(app *fiber.App, warehouseController *controllers.WarehouseController) {

	app.Post("/warehouses", middleware.AuthMiddleware, warehouseController.CreateWarehouse)
	app.Post("/warehouses/:warehouseID/shop/:shopID", middleware.AuthMiddleware, warehouseController.SetWarehouseShop)
	app.Get("/warehouses/:warehouseID/shops", middleware.AuthMiddleware, warehouseController.GetShopsByWarehouse)
	app.Get("/warehouses/shop/:shop_id", warehouseController.GetWarehousesByShopID) // New route
    app.Post("/warehouses/:warehouse_id/stock", warehouseController.AddStock)
    app.Post("/warehouses/transfer", warehouseController.TransferStock)
    app.Put("/warehouses/:warehouse_id/status", warehouseController.SetWarehouseStatus)
}