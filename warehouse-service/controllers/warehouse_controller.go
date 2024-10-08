package controllers

import (
	"strconv"

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

	warehouseIDInt, _ := strconv.ParseUint(warehouseID, 10, 32)
	
	_, err := ctrl.WarehouseService.GetWarehouseByID(uint(warehouseIDInt))
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

func (ctrl *WarehouseController) GetShopsByWarehouse(c *fiber.Ctx) error {
	warehouseID := c.Params("warehouseID")
	if warehouseID == "" {
		return utils.SendResponse(c, fiber.StatusBadRequest, "Warehouse ID is required", nil)	
	}	
	warehouseIDInt, _ := strconv.ParseUint(warehouseID, 10, 32)
	_, err := ctrl.WarehouseService.GetWarehouseByID(uint(warehouseIDInt))
	if err != nil {
		return utils.SendResponse(c, fiber.StatusNotFound, "Warehouse not found", nil)
	}

	warehouseShops, err := ctrl.WarehouseService.GetShopsWarehouse(warehouseID)
	if err != nil {
		return utils.SendResponse(c, fiber.StatusInternalServerError, "Failed to get warehouse shops", nil)
	}

	return utils.SendResponse(c, fiber.StatusOK, "Successfully get warehouse shops", warehouseShops)
}


// GetWarehousesByShopID handles fetching warehouses for a specific shop
func (ctrl *WarehouseController) GetWarehousesByShopID(c *fiber.Ctx) error {
    shopID := c.Params("shop_id") // Get shop ID from URL parameter
    id, err := strconv.ParseUint(shopID, 10, 32)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid shop ID"})
    }

    warehouses, err := ctrl.WarehouseService.GetWarehousesByShopID(uint(id))
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }

    return c.JSON(warehouses)
}

func (ctrl *WarehouseController) AddStock(c *fiber.Ctx) error {
    var request struct {
        WarehouseID uint `json:"warehouse_id"`
        ProductID   uint `json:"product_id"`
        Quantity    int  `json:"quantity"`
    }

    if err := c.BodyParser(&request); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
    }

    err := ctrl.WarehouseService.AddStock(request.WarehouseID, request.ProductID, request.Quantity)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Stock added successfully"})
}

func (ctrl *WarehouseController) TransferStock(c *fiber.Ctx) error {
    var request struct {
        FromWarehouseID uint `json:"from_warehouse_id"`
        ToWarehouseID   uint `json:"to_warehouse_id"`
        ProductID       uint `json:"product_id"`
        Quantity        int  `json:"quantity"`
    }

    if err := c.BodyParser(&request); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
    }

    err := ctrl.WarehouseService.TransferStock(request.FromWarehouseID, request.ToWarehouseID, request.ProductID, request.Quantity)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Stock transferred successfully"})
}

func (ctrl *WarehouseController) SetWarehouseStatus(c *fiber.Ctx) error {
    var request struct {
        IsActive bool `json:"is_active"`
    }

    warehouseID := c.Params("warehouseID")
	warehouseIDInt, _ := strconv.ParseUint(warehouseID, 10, 32)

    if err := c.BodyParser(&request); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
    }

    err := ctrl.WarehouseService.SetWarehouseStatus(uint(warehouseIDInt), request.IsActive)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Warehouse status updated successfully"})
}