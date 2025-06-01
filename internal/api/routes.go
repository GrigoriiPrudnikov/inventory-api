package api

import "github.com/gin-gonic/gin"

func SetupRoutes() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, "Hello World!")
	})

	r.GET("/users", GetUsers)
	r.GET("/users/:id", GetUser)
	r.POST("/users", CreateUser)

	r.GET("/items", GetItems)
	r.GET("/items/:id", GetItem)

	r.Run(":8080")
}

/*
POST /user - create user
POST /item - create item and add it to user (jwt)
POST /items/:id/sell - sell item
*/
