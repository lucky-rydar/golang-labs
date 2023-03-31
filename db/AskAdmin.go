package db

import (
	"time"
)

type AskAdmin struct {
	Id uint `gorm:"primaryKey;autoIncrement`
	Action string
	Name string
	Surname string
	IsMale  bool
	StudentTicketNumber string
	StudentTicketExpireDate time.Time
	RoomNumber string
}
