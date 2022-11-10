package service

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Init() {
	var err error
	Db, err = gorm.Open(sqlite.Open("database.sqlite"), &gorm.Config{})
	db, _ := Db.DB()
	db.SetConnMaxLifetime(1)
	if err != nil {
		panic(err)
	}
}
