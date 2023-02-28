package db

import (
	"fmt"
	"log"

	"github.com/it-02/dormitory/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database connected")

	err = DB.AutoMigrate(&models.Room{})
	if err != nil {
		log.Fatal(err)
	}

	println("Database initialized")
}
