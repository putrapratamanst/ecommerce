package controllers

import (
	"fmt"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/putrapratamanst/ecommerce/warehouse-service/models"
	"github.com/putrapratamanst/ecommerce/warehouse-service/services"
	"github.com/putrapratamanst/ecommerce/warehouse-service/utils"
)

type WarehouseController struct {
	WarehouseService *services.WarehouseService
	validate         *validator.Validate
	shopService      *services.ShopServiceClient
}

func NewWarehouseController(warehouseService *services.WarehouseService, shopService *services.ShopServiceClient) *WarehouseController {
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

	warehouseIDInt, _ := strconv.ParseUint(warehouseID, 10, 32)

	_, err := ctrl.WarehouseService.GetWarehouseByID(uint(warehouseIDInt))
	if err != nil {
		return utils.SendResponse(c, fiber.StatusNotFound, "Warehouse not found", nil)
	}

	_, errShop := ctrl.shopService.GetShopByID(shopID)
	if errShop != nil {
		return utils.SendResponse(c, fiber.StatusInternalServerError, "Failed to get shop details", nil)
	}

	errSet := ctrl.WarehouseService.SetWarehouseShop(warehouseID, shopID)
	if errSet != nil {
		fmt.Println(errSet)
		return utils.SendResponse(c, fiber.StatusInternalServerError, "Failed to set warehouse shop", nil)
	}

	return utils.SendResponse(c, fiber.StatusOK, "Successfully set warehouse shop", nil)
}

func (ctrl *WarehouseController) AdjustStock(c *fiber.Ctx) error {

	warehouseID, _ := c.ParamsInt("warehouseID")
	productID, _ := c.ParamsInt("productID")
	type request struct {
		Quantity int `json:"quantity" validate:"required"`
	}
	var req request

	if err := c.BodyParser(&req); err != nil {
		return utils.SendResponse(c, fiber.StatusBadRequest, "Invalid request body", nil)
	}

	if err := ctrl.validate.Struct(req); err != nil {
		return utils.SendResponse(c, fiber.StatusBadRequest, "Invalid input validation: "+err.Error(), nil)
	}

	if err := ctrl.WarehouseService.AdjustStock(warehouseID, productID, req.Quantity); err != nil {
		return utils.SendResponse(c, fiber.StatusInternalServerError, "Failed to adjust stock warehouse", nil)
	}

	return utils.SendResponse(c, fiber.StatusCreated, "Successfully adjust stock warehouse", nil)
}

func (ctrl *WarehouseController) TransferStock(c *fiber.Ctx) error {
	type request struct {
		FromWarehouseID int `json:"fromWarehouseID" validate:"required"`
		ToWarehouseID   int `json:"toWarehouseID" validate:"required"`
		ProductID       int `json:"productID" validate:"required"`
		Quantity        int `json:"quantity" validate:"required"`
	}
	var req request

	if err := c.BodyParser(&req); err != nil {
		return utils.SendResponse(c, fiber.StatusBadRequest, "Invalid request body", nil)
	}

	if err := ctrl.validate.Struct(req); err != nil {
		return utils.SendResponse(c, fiber.StatusBadRequest, "Invalid input validation: "+err.Error(), nil)
	}

	if err := ctrl.WarehouseService.TransferStock(req.FromWarehouseID, req.ToWarehouseID, req.ProductID, req.Quantity); err != nil {
		return utils.SendResponse(c, fiber.StatusInternalServerError, "Failed to transfer warehouse", nil)
	}

	return utils.SendResponse(c, fiber.StatusCreated, "Successfully transfer stock warehouse", nil)
}

func (ctrl *WarehouseController) ActivateWarehouse(c *fiber.Ctx) error {
	warehouseID, _ := c.ParamsInt("warehouseID")
	type request struct {
		Active bool `json:"active" validate:"required"`
	}
	var req request

	if err := c.BodyParser(&req); err != nil {
		return utils.SendResponse(c, fiber.StatusBadRequest, "Invalid request body", nil)
	}

	if err := ctrl.WarehouseService.SetWarehouseActive(warehouseID, req.Active); err != nil {
		return utils.SendResponse(c, fiber.StatusInternalServerError, "Failed to activate warehouse", nil)
	}

	return utils.SendResponse(c, fiber.StatusCreated, "Successfully activate warehouse", nil)
}
