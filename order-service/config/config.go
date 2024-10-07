package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/putrapratamanst/ecommerce/order-service/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var ctx = context.Background()

var JWT_SECRET = []byte(os.Getenv("JWT_SECRET"))

func InitDB() *gorm.DB {
	// Get DB connection details from environment variables
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbSSLMode := os.Getenv("DB_SSLMODE")

	// Build the connection string
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", dbHost, dbUser, dbPassword, dbName, dbPort, dbSSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	// Migrate the schema
	err = db.AutoMigrate(&models.Order{})
	if err != nil {
		log.Fatal("failed to migrate:", err)
	}

	log.Println("Migration successful!")

	return db
}

func LoadEnv() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
