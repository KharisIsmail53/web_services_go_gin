package main

import (
	"log"
	"gin-framework-services/models"
	"gin-framework-services/routes"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres dbname=masjid port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Migrate the schema
	db.AutoMigrate(&models.Barang{})

	// Initialize Gin router
	router := routes.APIRouter(db)

	// Setup routes
	

	// Run the server
	router.Run(":8080")
}
