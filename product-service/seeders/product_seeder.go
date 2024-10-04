package seeders

import (
	"github.com/putrapratamanst/ecommerce/product-service/models"
	"gorm.io/gorm"
)

func SeedProducts(db *gorm.DB) {
	products := []models.Product{
		{
			Name:        "Laptop Dell XPS 13",
			Price:       15000000,
			Description: "Laptop dengan desain premium dan performa tinggi, dilengkapi layar 13 inci.",
		},
		{
			Name:        "Apple MacBook Air M1",
			Price:       13000000,
			Description: "MacBook ringan dengan chip M1 yang memberikan performa luar biasa.",
		},
		{
			Name:        "Smartphone Samsung Galaxy S21",
			Price:       12000000,
			Description: "Smartphone flagship dengan kamera luar biasa dan tampilan yang menawan.",
		},
		{
			Name:        "Headphone Sony WH-1000XM4",
			Price:       4000000,
			Description: "Headphone noise-cancelling dengan kualitas suara yang sangat baik.",
		},
		{
			Name:        "Smartwatch Garmin Forerunner 245",
			Price:       5000000,
			Description: "Smartwatch untuk pelari dengan pelacakan GPS dan pemantauan detak jantung.",
		},
	}
	for _, product := range products {
		db.FirstOrCreate(&product, models.Product{Name: product.Name})
	}
}
