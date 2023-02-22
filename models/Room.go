package models

type Room struct {
	Id           uint `gorm:"primaryKey;autoIncrement"`
	IsMale       bool
	AreaSqMeters float32
}
