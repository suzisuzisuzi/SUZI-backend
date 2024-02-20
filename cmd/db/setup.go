package db

import (
	"fmt"
	"os"
	"os/user"

	"github.com/fauxriarty/SUZI-backend/cmd/data"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	host1 := os.Getenv("DBHOST")
	user1 := os.Getenv("DBUSER")
	password1 := os.Getenv("DBPASSWORD")
	dbname1 := os.Getenv("DB")
	port1 := os.Getenv("DBPORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", host1, user1, password1, dbname1, port1)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(user.User{}, data.Datalog{})
	if err != nil {
		return
	}
	DB = db
}
