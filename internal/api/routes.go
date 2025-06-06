package api

import (
	"inventory-api/internal/api/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) { c.JSON(200, "Hello World!") })

	r.POST("/auth/login", handlers.Login)
	r.POST("/auth/refresh", handlers.Refresh)

	r.GET("/users", handlers.GetUsers)
	r.GET("/users/:id", handlers.GetUser)
	r.POST("/users", handlers.CreateUser)

	r.GET("/items", handlers.GetItems)
	r.GET("/items/:id", handlers.GetItem)
	r.POST("/items", handlers.CreateItem)
	r.GET("/items/sell/:id", handlers.SellItem)

	r.Run(":8080")
}
