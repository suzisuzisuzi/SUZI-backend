package main

import (
	"github.com/fauxriarty/SUZI-backend/cmd/auth"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/user/:id", auth.GetUserByID)
	router.POST("/login", auth.Login)
	router.Run("localhost:8080")
}
