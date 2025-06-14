package handlers

import (
	"errors"
	"inventory-api/internal/db"
	"inventory-api/internal/models"
	"inventory-api/internal/utils"

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

type CreateItemReq struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Level       uint8  `json:"level"`
	Price       uint64 `json:"price"`
	OwnerID     uint   `json:"owner_id"`
}

func CreateItem(c *gin.Context) {
	database := db.DB

	var req CreateItemReq
	c.BindJSON(&req)

	var found models.Item
	res := database.First(&found, "name = ?", req.Name)

	if res.Error == nil {
		c.JSON(409, gin.H{
			"error": "item already exists",
		})
		return
	}

	item := models.Item{
		Name:        req.Name,
		Description: req.Description,
		Level:       req.Level,
		Price:       req.Price,
		OwnerID:     req.OwnerID,
	}
	res = database.Create(&item)
	if res.Error != nil {
		c.JSON(500, gin.H{
			"error": "failed to create item: " + res.Error.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"item": item})
}

func SellItem(c *gin.Context) {
	id := c.Param("id")
	database := db.DB

	var item models.Item
	res := database.First(&item, id)

	if res.Error != nil && errors.Is(res.Error, gorm.ErrRecordNotFound) {
		c.JSON(404, gin.H{
			"error": "item not found",
		})
		return
	}

	token := c.GetHeader("Authorization")

	claims, err := utils.ParseToken(token)
	if err != nil {
		c.JSON(401, gin.H{
			"error": "unauthorized",
		})
		return
	}

	userID := claims.(*utils.CustomClaims).ID
	if item.OwnerID != userID {
		c.JSON(401, gin.H{
			"error": "unauthorized",
		})
		return
	}

	user := models.User{
		ID: userID,
	}
	res = database.First(&user)
	if res.Error != nil {
		c.JSON(500, gin.H{
			"error": "failed to get user: " + res.Error.Error(),
		})
		return
	}

	user.Coins += item.Price
	res = database.Save(&user)
	if res.Error != nil {
		c.JSON(500, gin.H{
			"error": "failed to update user: " + res.Error.Error(),
		})
		return
	}

	res = database.Delete(&item)
	if res.Error != nil {
		c.JSON(500, gin.H{
			"error": "failed to delete item: " + res.Error.Error(),
		})
		return
	}

	c.JSON(200, "success")
}
