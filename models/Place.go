package models

type Place struct {
	Id           uint `gorm:"primaryKey;autoIncrement"`
	IsFree	     bool
	RoomId       uint
}
