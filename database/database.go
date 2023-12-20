package database

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"

	models "github.com/theartefak/artefak/app/Models"
)

var DB *gorm.DB

func InitDB() {
	var err error
	db, err := gorm.Open(sqlite.Open("database/database.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.User{})

	DB = db
}
