package auth

import (
	"net/http"

	"github.com/fauxriarty/SUZI-backend/cmd/user"

	"github.com/fauxriarty/SUZI-backend/cmd/db"
	"github.com/gin-gonic/gin"
)

func GetUserByID(c *gin.Context) {
	var user user.User
	firebaseID := c.Param("firebaseID") // this should match the parameter name in the route

	if firebaseID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Firebase ID is required"})
		return
	}

	if err := db.DB.First(&user, "firebase_id = ?", firebaseID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func Login(c *gin.Context) {
	var user user.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.First(&user, "firebaseID = ?", user.FirebaseID).Error; err != nil {
		db.DB.Create(&user)
	}

	c.JSON(http.StatusOK, user)
}
