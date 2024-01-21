package user

type User struct {
	// ID         uint   `json:"id" ` //id has to be an unsigned integer
	Username   string `json:"username"`
	FirebaseID string `json:"firebaseID" gorm:"primaryKey"`
	State      string `json:"state"`
}
