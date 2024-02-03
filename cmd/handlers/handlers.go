package handlers

import (
	"net/http"

	"github.com/fauxriarty/SUZI-backend/cmd/data"
	"github.com/fauxriarty/SUZI-backend/cmd/db"
	"github.com/gin-gonic/gin"
)

func LogEntry(c *gin.Context) {
	firebaseID := c.Param("firebaseID") // this should match the parameter name in the route

	if firebaseID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Firebase ID is required"})
		return
	}

	var datalog data.Datalog
	if err := c.ShouldBindJSON(&datalog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	datalog.FirebaseID = firebaseID
	db.DB.Create(&datalog)

	c.JSON(http.StatusOK, datalog)
}

func GetHeatmapData(c *gin.Context) {
	category := c.Param("Category")

	if category == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Category is required"})
		return
	}

	var datalogs []data.Datalog
	db.DB.Where("category = ?", category).Find(&datalogs)

	var geojsons []data.GeoJSON
	for _, datalog := range datalogs {
		geojsons = append(geojsons, data.GeoJSON{
			Type:       "Feature",
			Properties: datalog,
			Geometry: data.Geometry{
				Type:        "Point",
				Coordinates: []float64{datalog.Longitude, datalog.Latitude},
			},
		})
	}

	c.JSON(http.StatusOK, geojsons)
}
