package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

//have to connect to database
var users = []User{
	{ID: "22bds0179", Username: "Aditya"},
	{ID: "22bee0014", Username: "Chandram"},
}

func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	for _, user := range users {
		if user.ID == id {
			c.JSON(http.StatusOK, user)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
}

func Login(c *gin.Context) {
	var newUser User

	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the user already exists
	for _, user := range users {
		if user.Username == newUser.Username {
			c.JSON(http.StatusConflict, gin.H{"message": "username already exists"})
			return
		}
	}

	// Append the new user to the users slice
	users = append(users, newUser)

	c.JSON(http.StatusOK, gin.H{"message": "user registered successfully", "user": newUser})
}
