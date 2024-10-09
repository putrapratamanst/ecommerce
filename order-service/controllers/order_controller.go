package controllers

import (
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/putrapratamanst/ecommerce/order-service/models"
	"github.com/putrapratamanst/ecommerce/order-service/services"
	"github.com/putrapratamanst/ecommerce/order-service/utils"
)

type OrderController struct {
	orderService   *services.OrderService
	validate       *validator.Validate
	productService *services.ProductServiceClient
}

func NewOrderController(orderService *services.OrderService, productService *services.ProductServiceClient) *OrderController {
	return &OrderController{orderService: orderService,
		validate:       validator.New(),
		productService: productService,
	}
}

func (oc *OrderController) CheckoutOrder(c *fiber.Ctx) error {
	uID := c.Locals("userID")
	userID, _ := strconv.Atoi(uID.(string))

	var order models.Order

	if err := c.BodyParser(&order); err != nil {
		return utils.SendResponse(c, fiber.StatusBadRequest, err.Error(), nil)
	}

	order.UserID = uint(userID)
	if err := oc.validate.Struct(order); err != nil {
		return utils.SendResponse(c, fiber.StatusBadRequest, "Invalid input validation: "+err.Error(), nil)
	}

	product, errProduct := oc.productService.GetProductByID(order.ProductID)
	if errProduct != nil {
		return utils.SendResponse(c, fiber.StatusInternalServerError, "Failed to get product details", nil)
	}

	order.TotalPrice = float64(order.Quantity) * product.Data.Price
	err := oc.orderService.CheckoutOrder(c.Context(), &order)
	if err != nil {
		return utils.SendResponse(c, fiber.StatusInternalServerError, err.Error(), nil)
	}

	return utils.SendResponse(c, fiber.StatusOK, "Successfully checkout order", order)
}

func (oc *OrderController) PaymentConfirm(c *fiber.Ctx) error {
	var paymentInfo struct {
		OrderID int `validate:"required"`
	}

	if err := c.BodyParser(&paymentInfo); err != nil {
		return utils.SendResponse(c, fiber.StatusBadRequest, err.Error(), nil)
	}

	if err := oc.validate.Struct(paymentInfo); err != nil {
		return utils.SendResponse(c, fiber.StatusBadRequest, "Invalid input validation: "+err.Error(), nil)
	}
	order, err := oc.orderService.GetOrderByID(paymentInfo.OrderID)
	if err != nil {
		return utils.SendResponse(c, fiber.StatusInternalServerError, err.Error(), nil)
	}

	order.Status = "PAID"
	errUpdate := oc.orderService.UpdateOrderStatus(order)
	if errUpdate != nil {
		return utils.SendResponse(c, fiber.StatusInternalServerError, errUpdate.Error(), nil)
	}

	return utils.SendResponse(c, fiber.StatusOK, "Successfully confirm payment", nil)
}

func (oc *OrderController) CancelOrder(c *fiber.Ctx) error {
	var paymentInfo struct {
		OrderID int `validate:"required"`
	}

	if err := c.BodyParser(&paymentInfo); err != nil {
		return utils.SendResponse(c, fiber.StatusBadRequest, err.Error(), nil)
	}

	if err := oc.validate.Struct(paymentInfo); err != nil {
		return utils.SendResponse(c, fiber.StatusBadRequest, "Invalid input validation: "+err.Error(), nil)
	}

	order, err := oc.orderService.GetOrderByID(paymentInfo.OrderID)
	if err != nil {
		return utils.SendResponse(c, fiber.StatusInternalServerError, err.Error(), nil)
	}

	if order.Status == "PAID" {
		return utils.SendResponse(c, fiber.StatusOK, "Paid order cannot be canceled", nil)
	}

	// Jika order belum dibayar, kita akan mengirim pesan untuk melepaskan stok
	err = oc.orderService.ReleaseOrder(c.Context(), order)
	if err != nil {
		return err
	}

	return utils.SendResponse(c, fiber.StatusOK, "Successfully cancel order", nil)
}
