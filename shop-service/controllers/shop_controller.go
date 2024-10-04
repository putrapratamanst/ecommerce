package controllers

import (
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/putrapratamanst/ecommerce/shop-service/models"
	"github.com/putrapratamanst/ecommerce/shop-service/services"
	"github.com/putrapratamanst/ecommerce/shop-service/utils"
)

type ShopController struct {
	ShopService services.ShopService
	validate    *validator.Validate
}

func NewShopController(shopService services.ShopService) *ShopController {
	return &ShopController{
		ShopService: shopService,
		validate:    validator.New(),
	}
}

func (ctrl *ShopController) CreateShop(c *fiber.Ctx) error {
    userID := c.Locals("userID")
	ownerID, _ := strconv.Atoi(userID.(string))
	
	var shop models.Shop
	if err := c.BodyParser(&shop); err != nil {
		return utils.SendResponse(c, fiber.StatusBadRequest, "Invalid request body", nil)
	}

	shop.OwnerID = ownerID
	if err := ctrl.validate.Struct(shop); err != nil {
		return utils.SendResponse(c, fiber.StatusBadRequest, "Invalid input validation: "+err.Error(), nil)
	}

	if err := ctrl.ShopService.CreateShop(&shop); err != nil {
		return utils.SendResponse(c, fiber.StatusInternalServerError, "Failed to create shop", nil)
	}

	return utils.SendResponse(c, fiber.StatusCreated, "Successfully created shop", shop)
}

func (ctrl *ShopController) GetShop(c *fiber.Ctx) error {
	shopID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.SendResponse(c, fiber.StatusBadRequest, "Invalid shop ID", nil)
	}

	shop, err := ctrl.ShopService.GetShopByID(shopID)
	if err != nil {
		return utils.SendResponse(c, fiber.StatusNotFound, "Shop not found", nil)
	}

	return utils.SendResponse(c, fiber.StatusOK, "Successfully retrieved shop details", shop)
}

func (ctrl *ShopController) UpdateShop(c *fiber.Ctx) error {
	shopID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.SendResponse(c, fiber.StatusBadRequest, "Invalid shop ID", nil)
	}

	var shop models.Shop
	if err := c.BodyParser(&shop); err != nil {
		return utils.SendResponse(c, fiber.StatusBadRequest, "Invalid request body", nil)
	}

	shop.ID = uint(shopID)

	if err := ctrl.ShopService.UpdateShop(&shop); err != nil {
		return utils.SendResponse(c, fiber.StatusNotFound, "Shop not found", nil)
	}

	return utils.SendResponse(c, fiber.StatusOK, "Successfully updated shop", shop)
}

func (ctrl *ShopController) DeleteShop(c *fiber.Ctx) error {
	shopID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.SendResponse(c, fiber.StatusBadRequest, "Invalid shop ID", nil)
	}

	if err := ctrl.ShopService.DeleteShop(shopID); err != nil {
		return utils.SendResponse(c, fiber.StatusNotFound, "Shop not found", nil)
	}

	return utils.SendResponse(c, fiber.StatusNoContent, "Successfully deleted shop", nil) // 204 No Content
}

func (ctrl *ShopController) GetAllShops(c *fiber.Ctx) error {
	shops, err := ctrl.ShopService.GetAllShops()
	if err != nil {
		return utils.SendResponse(c, fiber.StatusInternalServerError, "Failed to retrieve shops", nil)
	}

	return utils.SendResponse(c, fiber.StatusOK, "Successfully retrieved shops", shops)
}
