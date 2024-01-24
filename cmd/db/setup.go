package db

import (
	"github.com/fauxriarty/SUZI-backend/cmd/data"
	"github.com/fauxriarty/SUZI-backend/cmd/user"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=db.lqdkfanqochmikddbugv.supabase.co user=postgres password=Thepunisintheoven@1 dbname=postgres port=5432"
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
