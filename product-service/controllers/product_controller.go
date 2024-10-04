package controllers

import (
	"strconv"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/putrapratamanst/ecommerce/product-service/repositories"
	"github.com/putrapratamanst/ecommerce/product-service/services"
	"github.com/putrapratamanst/ecommerce/product-service/utils"
	"gorm.io/gorm"
)

type ProductController struct {
	ProductService *services.ProductService
}

func NewProductController(db *gorm.DB, redisClient *redis.Client) *ProductController {
	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo, redisClient)
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
