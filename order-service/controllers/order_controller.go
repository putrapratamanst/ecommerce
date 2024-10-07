package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/putrapratamanst/ecommerce/order-service/services"
	"github.com/putrapratamanst/ecommerce/order-service/utils"
)

type OrderController struct {
	orderService services.OrderService
}

func NewOrderController(orderService services.OrderService) *OrderController {
	return &OrderController{orderService: orderService}
}

func (oc *OrderController) PlaceOrder(c *fiber.Ctx) error {
    userID := c.Locals("userID")

	var request struct {
		ProductID uint `json:"product_id"`
		Quantity  int  `json:"quantity"`
	}

	if err := c.BodyParser(&request); err != nil {
		return utils.SendResponse(c, fiber.StatusBadRequest, err.Error(), nil)
	}

	order, err := oc.orderService.PlaceOrder(userID.(uint), request.ProductID, request.Quantity)
	if err != nil {
		return utils.SendResponse(c, fiber.StatusInternalServerError, err.Error(), nil)
	}

	return c.Status(fiber.StatusCreated).JSON(order)
}

func (oc *OrderController) GetOrder(c *fiber.Ctx) error {
	orderID := c.Params("id")
    u64, _ := strconv.ParseUint(orderID, 10, 32)
   
	order, err := oc.orderService.GetOrderByID(uint(u64))
	if err != nil {
		return utils.SendResponse(c, fiber.StatusNotFound, "Order not found", nil)
	}

	return c.JSON(order)
}

func (oc *OrderController) GetOrders(c *fiber.Ctx) error {
    userID := c.Locals("userID")

	orders, err := oc.orderService.GetOrdersByUserID(userID.(uint))
	if err != nil {
		return utils.SendResponse(c, fiber.StatusInternalServerError, err.Error(), nil)
	}

	return c.JSON(orders)
}
