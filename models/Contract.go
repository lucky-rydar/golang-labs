package models

import (
	"time"
)

type Contract struct {
	Id         uint `gorm:"primaryKey;autoIncrement"`
	SignDate   time.Time
	ExpireDate time.Time
}
