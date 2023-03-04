package models

import (
	"time"
)

type StudentTicket struct {
	Id           uint `gorm:"primaryKey;autoIncrement"`
	SerialNumber string
	ExpireDate   time.Time
}
