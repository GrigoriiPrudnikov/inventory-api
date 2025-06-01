package handlers

import (
	"errors"
	"inventory-api/internal/db"
	"inventory-api/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetItems(c *gin.Context) {
	database := db.DB

	var items []models.Item
	database.Find(&items)

	c.JSON(200, gin.H{"items": items})
}

func GetItem(c *gin.Context) {
	id := c.Param("id")
	database := db.DB

	var item models.Item
	res := database.First(&item, id)

	if res.Error != nil && errors.Is(res.Error, gorm.ErrRecordNotFound) {
		c.JSON(404, gin.H{
			"item":  nil,
			"error": "item not found",
		})
		return
	}

	c.JSON(200, gin.H{"item": item})
}
