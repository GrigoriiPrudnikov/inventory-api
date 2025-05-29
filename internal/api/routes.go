package api

import "github.com/gin-gonic/gin"

func SetupRoutes() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, "Hello World!")
	})
	r.Run(":8080")
}

/*
user: username, password, items, coins
item: name, description, price, level, quantity

ENDPOINTS:

POST /user - create user
POST /item - create item and add it to user (jwt)
POST /items/:id/sell - sell item
*/
