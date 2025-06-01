package api

import (
	"inventory-api/internal/api/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, "Hello World!")
	})

	r.GET("/users", handlers.GetUsers)
	r.GET("/users/:id", handlers.GetUser)
	r.POST("/users", handlers.CreateUser)

	r.GET("/items", handlers.GetItems)
	r.GET("/items/:id", handlers.GetItem)

	r.Run(":8080")
}

/*
POST /item - create item and add it to user (jwt)
POST /items/:id/sell - sell item
*/
