package db

import (
	"github.com/fauxriarty/SUZI-backend/cmd/user"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=postgres.suzi-backend.orb.local user=myuser password=mypassword dbname=mydatabase port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(user.User{})
	if err != nil {
		return
	}
	DB = db
}
