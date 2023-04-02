package db

type Room struct {
	Id           uint `gorm:"primaryKey;autoIncrement"`
	Number	     string // number is not only an integer, but also a letter
	IsMale       bool
	AreaSqMeters float32
}
