package auth

import (
	"github.com/gin-gonic/gin"
)

type User struct {
	Username  string `json:"username"`
	AuthToken string `json:"authtoken"`
}

var users = []User{}

func getUserByID(c *gin.Context) {

}

func login(c *gin.Context) {

}

func main() {
	router := gin.Default()
	router.GET("/user:id", getUserByID)
	router.POST("/login", login)

	router.Run("localhost:8080")
}
