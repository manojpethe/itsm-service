package repository

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

var Databasefile = "itsm.db"

func Connect() {

	var errDBconnect error
	DB, errDBconnect = gorm.Open(sqlite.Open(Databasefile), &gorm.Config{})
	if errDBconnect != nil {
		panic("failed to connect database")
	}
}
