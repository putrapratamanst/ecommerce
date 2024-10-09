package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/putrapratamanst/ecommerce/order-service/models"
	"github.com/putrapratamanst/ecommerce/order-service/services"
	"github.com/putrapratamanst/ecommerce/order-service/utils"
)

type OrderController struct {
	orderService *services.OrderService
}

func NewOrderController(orderService *services.OrderService) *OrderController {
	return &OrderController{orderService: orderService}
}

func (oc *OrderController) CheckoutOrder(c *fiber.Ctx) error {
    userID := c.Locals("userID")
	var order models.Order


	if err := c.BodyParser(&order); err != nil {
		return utils.SendResponse(c, fiber.StatusBadRequest, err.Error(), nil)
	}

	order.UserID = userID.(uint)
	err := oc.orderService.CheckoutOrder(c.Context(), &order)
	if err != nil {
		return utils.SendResponse(c, fiber.StatusInternalServerError, err.Error(), nil)
	}

	return utils.SendResponse(c, fiber.StatusOK, "Successfully checkout order", order)
}

func (oc *OrderController) PaymentConfirm(c *fiber.Ctx) error {
	var paymentInfo struct {
		OrderID int `json:"order_id" validate:"required"`
	}
	
	if err := c.BodyParser(&paymentInfo); err != nil {
		return utils.SendResponse(c, fiber.StatusBadRequest, err.Error(), nil)
	}

	order, err := oc.orderService.GetOrderByID(paymentInfo.OrderID)
	if err != nil {
		return utils.SendResponse(c, fiber.StatusInternalServerError, err.Error(), nil)
	}

	errUpdate := oc.orderService.UpdateOrderStatus(order)
	if errUpdate != nil {
		return utils.SendResponse(c, fiber.StatusInternalServerError, errUpdate.Error(), nil)
	}

	return utils.SendResponse(c, fiber.StatusOK, "Successfully confirm payment", order)
}

func (oc *OrderController) CancelOrder(c *fiber.Ctx) error {
	var paymentInfo struct {
		OrderID int `json:"order_id" validate:"required"`
	}
	
	if err := c.BodyParser(&paymentInfo); err != nil {
		return utils.SendResponse(c, fiber.StatusBadRequest, err.Error(), nil)
	}

	order, err := oc.orderService.GetOrderByID(paymentInfo.OrderID)
	if err != nil {
		return utils.SendResponse(c, fiber.StatusInternalServerError, err.Error(), nil)
	}

	// Jika order belum dibayar, kita akan mengirim pesan untuk melepaskan stok
    if order.Status != "PAID" {
        err := oc.orderService.ReleaseOrder(c.Context(), order)
        if err != nil {
            return err
        }
    }

	order.Status = "CANCELLED"
	errUpdate := oc.orderService.UpdateOrderStatus(order)
	if errUpdate != nil {
		return utils.SendResponse(c, fiber.StatusInternalServerError, errUpdate.Error(), nil)
	}

	return utils.SendResponse(c, fiber.StatusOK, "Successfully cancel order", order)
}

