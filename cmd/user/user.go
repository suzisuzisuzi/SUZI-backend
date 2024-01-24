package user

import "github.com/fauxriarty/SUZI-backend/cmd/data"

type User struct {
	// ID         uint   `json:"id" ` //id has to be an unsigned integer
	Email      string `json:"email"`
	FirebaseID string `json:"firebaseID" gorm:"primaryKey"`
	State      string `json:"state"`

	Datalogs []data.Datalog `gorm:"foreignKey:FirebaseID"` // Use the fully qualified type name
}
