package config

import (
	"log"

	"github.com/putrapratamanst/ecommerce/product-service/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=products port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	
	// Migrate the schema
	err = db.AutoMigrate(&models.Product{})
	if err != nil {
		log.Fatal("failed to migrate:", err)
	}

	log.Println("Migration successful!")
	return db
}