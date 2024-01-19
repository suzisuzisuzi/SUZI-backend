package main

import (
	"net/http"
	"github.com/fauxriarty/SUZI-backend/cmd/auth"
	"github.com/gin-gonic/gin"
)

func Status(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func main() {
	router := gin.Default()
	router.GET("/", Status)
	router.GET("/user/:id", auth.GetUserByID)
	router.POST("/login", auth.Login)
	http.ListenAndServe(":8080", router)
}
