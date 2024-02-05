package data

import (
	"time"
)

type Datalog struct {
	Timestamp  time.Time `json:"timestamp"`
	FirebaseID string    `json:"firebaseID" gorm:"index"`
	Latitude   float64   `json:"latitude"`
	Longitude  float64   `json:"longitude"`
	Altitude   float64   `json:"altitude"`
	Category   string    `json:"category"`
}

type GeoJSON struct {
	Type       string   `json:"type"`
	Properties Datalog  `json:"properties"`
	Geometry   Geometry `json:"geometry"`
}

type Geometry struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}
