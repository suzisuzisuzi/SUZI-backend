package auth

import (
	"net/http"

	"github.com/fauxriarty/SUZI-backend/cmd/user"

	"github.com/fauxriarty/SUZI-backend/cmd/db"
	"github.com/gin-gonic/gin"
)

func GetUserByID(c *gin.Context) {
	id := c.Param("id")

	var newUser user.User
	db.DB.First(&newUser, id)

	if newUser.ID >= 0 {
		c.JSON(http.StatusOK, gin.H{"message": "user found", "user": newUser})
		return
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
}

func Login(c *gin.Context) {
	var user user.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User saved successfully"})
}
