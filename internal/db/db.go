package db

import (
	"database/sql"
	"inventory-api/config"
	"inventory-api/internal/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *sql.DB

func Init(cfg config.Config) {
	db, err := gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	err = db.AutoMigrate(&models.Item{}, &models.User{})
	if err != nil {
		log.Fatal("failed to auto migrate database")
	}

	postgresDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get generic database object: %v", err)
	}

	if err = postgresDB.Ping(); err != nil {
		log.Fatalf("database ping failed: %v", err)
	}
	log.Println("Database connection successful.")

	DB = postgresDB
}
