package db

import (
	"inventory-api/config"
	"inventory-api/internal/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init(cfg config.Config) {
	db, err := gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	err = db.AutoMigrate(&models.Item{}, &models.User{})
	if err != nil {
		log.Fatal("failed to auto migrate database")
	}

	connection, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get database connection: %v", err)
	}
	if err = connection.Ping(); err != nil {
		log.Fatalf("database ping failed: %v", err)
	}
	log.Println("Database connection successful.")

	DB = db
}
