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

	if newUser.ID != "" {
		c.JSON(http.StatusOK, gin.H{"message": "user found", "user": newUser})
		return
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
}

func Login(c *gin.Context) {
	var newUser user.User

	c.JSON(http.StatusOK, gin.H{"message": "user registered successfully", "user": newUser})
}
