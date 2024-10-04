package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/putrapratamanst/ecommerce/product-service/services"
	"github.com/putrapratamanst/ecommerce/product-service/utils"
)

type ProductController struct {
	ProductService services.ProductService
}

func NewProductController(productService services.ProductService) *ProductController {
	return &ProductController{ProductService: productService}
}

func (ctrl *ProductController) GetProducts(ctx *fiber.Ctx) error {
	page, _ := strconv.Atoi(ctx.Query("page", "1"))
	limit, _ := strconv.Atoi(ctx.Query("limit", "10"))

	offset := (page - 1) * limit

	products, total, err := ctrl.ProductService.GetProducts(limit, offset)
	if err != nil {
		return utils.SendResponse(ctx, fiber.StatusInternalServerError, "Cannot fetch products", nil)
	}

	pagination := utils.Paginate(total, page, limit)

	return utils.SendResponse(ctx, fiber.StatusOK, "Successfully fetched products", fiber.Map{
		"data":       products,
		"pagination": pagination,
	})
}
