package api

import (
	"errors"
	"inventory-api/internal/db"
	"inventory-api/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetUsers(c *gin.Context) {
	database := db.DB

	var users []models.User
	database.Find(&users)

	c.JSON(200, gin.H{"users": users})
}

func GetUser(c *gin.Context) {
	id := c.Param("id")
	database := db.DB

	var user models.User
	res := database.First(&user, id)

	if res.Error != nil && errors.Is(res.Error, gorm.ErrRecordNotFound) {
		c.JSON(404, gin.H{
			"user":  nil,
			"error": "user not found",
		})
		return
	}

	c.JSON(200, gin.H{"user": user})
}
