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
	Rating     float64   `json:"rating"`
}
