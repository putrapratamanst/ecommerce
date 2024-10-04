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
}

func NewWarehouseController(warehouseService services.WarehouseService) *WarehouseController {
	return &WarehouseController{
		WarehouseService: warehouseService,
		validate:         validator.New(),
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
