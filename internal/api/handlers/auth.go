package handlers

import (
	"errors"
	"inventory-api/internal/db"
	"inventory-api/internal/models"
	"inventory-api/internal/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var req LoginReq
	c.BindJSON(&req)

	database := db.DB

	var user models.User
	res := database.First(&user, "username = ?", req.Username)

	if res.Error != nil && errors.Is(res.Error, gorm.ErrRecordNotFound) {
		c.JSON(404, gin.H{
			"error": "user not found",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		c.JSON(401, gin.H{
			"error": "invalid username or password",
		})
		return
	}

	token := utils.GenerateToken(jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour).Unix(),
	})
	refresh := utils.GenerateToken(jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24 * 7).Unix(),
	})

	if token == nil || refresh == nil {
		c.JSON(500, gin.H{
			"error": "failed to sign token",
		})
		return
	}

	c.JSON(200, gin.H{
		"token":        token,
		"refreshToken": refresh,
	})
}
