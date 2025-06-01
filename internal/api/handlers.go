package api

import (
	"errors"
	"inventory-api/internal/db"
	"inventory-api/internal/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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
			"error": "user not found",
		})
		return
	}

	c.JSON(200, gin.H{"user": user})
}

type CreateUserReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func CreateUser(c *gin.Context) {
	database := db.DB

	var req CreateUserReq
	c.BindJSON(&req)

	var found models.User
	res := database.First(&found, "username = ?", req.Username)

	if res.Error == nil {
		c.JSON(409, gin.H{
			"error": "user already exists",
		})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "failed to hash password",
		})
		return
	}

	user := models.User{
		Username: req.Username,
		Password: string(hashedPassword),
	}
	res = database.Create(&user)
	if res.Error != nil {
		c.JSON(500, gin.H{
			"error": "failed to create user: " + res.Error.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"user": user})
}

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
