package user

type User struct {
	ID         uint   `json:"id" gorm:"primaryKey"` //id has to be an unsigned integer
	Username   string `json:"username"`
	FirebaseID string `json:"firebaseID"`
}
