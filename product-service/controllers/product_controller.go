package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/putrapratamanst/ecommerce/product-service/repositories"
	"github.com/putrapratamanst/ecommerce/product-service/utils"
	"github.com/putrapratamanst/ecommerce/product-service/services"
	"gorm.io/gorm"
)

type ProductController struct {
    ProductService *services.ProductService
}

func NewProductController(db *gorm.DB) *ProductController {
    productRepo := repositories.NewProductRepository(db)
    productService := services.NewProductService(productRepo)
    return &ProductController{ProductService: productService}
}

func (c *ProductController) GetProducts(ctx *fiber.Ctx) error {
    page, _ := strconv.Atoi(ctx.Query("page", "1"))
    limit, _ := strconv.Atoi(ctx.Query("limit", "10"))

    offset := (page - 1) * limit

    products, total, err := c.ProductService.GetProducts(limit, offset)
    if err != nil {
        return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot fetch products"})
    }

    pagination := utils.Paginate(total, page, limit)

    return ctx.JSON(fiber.Map{
        "data":       products,
        "pagination": pagination,
    })
}
