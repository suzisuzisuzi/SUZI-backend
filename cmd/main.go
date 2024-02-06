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

	router.Use(corsMiddleware())

	db.ConnectDatabase()

	router.GET("/", Status)
	router.GET("/user/:firebaseID", auth.GetUserByID)
	router.POST("/login", auth.Login)

	router.POST("/entries/:firebaseID", handlers.LogEntry)
	router.GET("/heatmap/:Category", handlers.GetHeatmapData)

	router.GET("/gheatmap/:Category", handlers.GetGheatmapData)

	http.ListenAndServe(":8080", router)

}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length, Content-Range")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
