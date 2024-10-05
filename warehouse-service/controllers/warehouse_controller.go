package controllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/putrapratamanst/ecommerce/warehouse-service/models"
	"github.com/putrapratamanst/ecommerce/warehouse-service/services"
	"github.com/putrapratamanst/ecommerce/warehouse-service/utils"
)

type WarehouseController struct {
	WarehouseService services.WarehouseService
	validate         *validator.Validate
	shopService      *services.ShopServiceClient
}

func NewWarehouseController(warehouseService services.WarehouseService, shopService *services.ShopServiceClient) *WarehouseController {
	return &WarehouseController{
		WarehouseService: warehouseService,
		validate:         validator.New(),
		shopService:      shopService,
	}
}

func (ctrl *WarehouseController) CreateWarehouse(c *fiber.Ctx) error {
	var warehouse models.Warehouse
	if err := c.BodyParser(&warehouse); err != nil {
		return utils.SendResponse(c, fiber.StatusBadRequest, "Invalid request body", nil)
	}

	if err := ctrl.validate.Struct(warehouse); err != nil {
		return utils.SendResponse(c, fiber.StatusBadRequest, "Invalid input validation: "+err.Error(), nil)
	}

	if err := ctrl.WarehouseService.CreateWarehouse(&warehouse); err != nil {
		return utils.SendResponse(c, fiber.StatusInternalServerError, "Failed to create warehouse", nil)
	}

	return utils.SendResponse(c, fiber.StatusCreated, "Successfully created warehouse", warehouse)
}

func (ctrl *WarehouseController) SetWarehouseShop(c *fiber.Ctx) error {
	warehouseID := c.Params("warehouseID")
	shopID := c.Params("shopID")

	if warehouseID == "" || shopID == "" {
		return utils.SendResponse(c, fiber.StatusBadRequest, "Warehouse ID and Shop ID are required", nil)
	}

	_, errCheckWarehouse := ctrl.WarehouseService.GetWarehouseShop(warehouseID, shopID)
	if errCheckWarehouse == nil {
		return utils.SendResponse(c, fiber.StatusInternalServerError, "Successfully set warehouse shop", nil)
	}
	
	_, err := ctrl.WarehouseService.GetWarehouseByID(warehouseID)
	if err != nil {
		return utils.SendResponse(c, fiber.StatusNotFound, "Warehouse not found", nil)
	}

	_, errShop := ctrl.shopService.GetShopByID(shopID)
	if errShop != nil {
		return utils.SendResponse(c, fiber.StatusInternalServerError, "Failed to get shop details", nil)
	}

	err = ctrl.WarehouseService.SetWarehouseShop(warehouseID, shopID)
	if err != nil {
		return utils.SendResponse(c, fiber.StatusInternalServerError, "Failed to set warehouse shop", nil)
	}

	return utils.SendResponse(c, fiber.StatusOK, "Successfully set warehouse shop", nil)
}
