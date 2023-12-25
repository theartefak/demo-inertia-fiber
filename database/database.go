package database

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

func New() (*Database, error) {
	var db *gorm.DB
	var err error

	db, err = gorm.Open(sqlite.Open("database/database.db"), &gorm.Config{})

	return &Database{db}, err
}
