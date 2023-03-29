package db

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// return db
func InitDB() *gorm.DB { 
	var err error
	var db *gorm.DB

	db, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database connected")

	err = db.AutoMigrate(
		&Room{},
		&Place{},
		&Contract{},
		&StudentTicket{},
		&Student{})

	if err != nil {
		log.Fatal(err)
	}

	println("Database initialized")
	return db
}

func SetupDB() {
	DB = InitDB()
}
