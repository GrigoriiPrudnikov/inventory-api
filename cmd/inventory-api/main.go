package main

import (
	"inventory-api/config"
	"inventory-api/internal/api"
	"inventory-api/internal/db"
)

func main() {
	config := config.Load()
	db.Init(config)
	api.SetupRoutes()
}
