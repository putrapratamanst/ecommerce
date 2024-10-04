package routes

import (
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/putrapratamanst/ecommerce/product-service/controllers"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB, redisClient *redis.Client) {
	// Initialize Product controller
	productController := controllers.NewProductController(db, redisClient)

	// example route: http://127.0.0.1:3000/products?page=2&limit=2
	app.Get("/products", productController.GetProducts)

}
