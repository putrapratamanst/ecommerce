package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/putrapratamanst/ecommerce/warehouse-service/controllers"
)

func SetupWarehouseRoutes(app *fiber.App, warehouseController *controllers.WarehouseController) {

	app.Post("/warehouse", warehouseController.CreateWarehouse)
}