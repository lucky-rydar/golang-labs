package db

type User struct {
	Id       uint   `gorm:"primaryKey;autoIncrement"`
	UUID     string `gorm:"uniqueIndex"`
	Username string
	Password string
	IsAdmin  bool
}
