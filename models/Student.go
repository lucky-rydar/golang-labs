package models

type Student struct {
	Id              uint `gorm:"primaryKey;autoIncrement"`
	Name            string
	Surname         string
	IsMale          bool
	ContractId      uint
	StudentTicketId uint
	PlaceId         uint
}
