package user

type User struct {
	ID       string `json:"id" gorm:"primaryKey"`
	Username string `json:"username"`
}
