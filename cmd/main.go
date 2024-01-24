package main

import (
	"net/http"

	"github.com/fauxriarty/SUZI-backend/cmd/auth"
	"github.com/fauxriarty/SUZI-backend/cmd/db"
	"github.com/fauxriarty/SUZI-backend/cmd/handlers"
	"github.com/gin-gonic/gin"
)

func Status(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func main() {
	router := gin.Default()

	db.ConnectDatabase()

	router.GET("/", Status)
	router.GET("/user/:firebaseID", auth.GetUserByID)
	router.POST("/login", auth.Login)

	router.POST("/entries/:firebaseID", handlers.LogEntry)
	router.GET("/heatmap/:Category", handlers.GetHeatmapData)

	http.ListenAndServe(":8080", router)

}
