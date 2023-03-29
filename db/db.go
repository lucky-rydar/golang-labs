package db

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Contract struct {
	Id         uint `gorm:"primaryKey;autoIncrement"`
	SignDate   time.Time
	ExpireDate time.Time
}

type Place struct {
	Id           uint `gorm:"primaryKey;autoIncrement"`
	IsFree	     bool
	RoomId       uint
}

type Room struct {
	Id           uint `gorm:"primaryKey;autoIncrement"`
	Number	     string // number is not only an integer, but also a letter
	IsMale       bool
	AreaSqMeters float32
}

type Student struct {
	Id              uint `gorm:"primaryKey;autoIncrement"`
	Name            string
	Surname         string
	IsMale          bool
	ContractId      uint
	StudentTicketId uint
	PlaceId         uint
}

type StudentTicket struct {
	Id           uint `gorm:"primaryKey;autoIncrement"`
	SerialNumber string
	ExpireDate   time.Time
}

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
