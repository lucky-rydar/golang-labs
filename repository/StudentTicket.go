package repository

import (
	"fmt"

	"github.com/it-02/dormitory/db"
	"gorm.io/gorm"
)

type IStudentTicket interface {
	AddStudentTicket(ticket *db.StudentTicket) error
	GetStudentTickets() []db.StudentTicket
	GetStudentTicketBySerialNumber(serialNumber string) db.StudentTicket
	GetStudentTicketById(id uint) db.StudentTicket
}

type StudentTicket struct {
	db *gorm.DB
}

func NewStudentTicket(db *gorm.DB) IStudentTicket {
	return &StudentTicket{db: db}
}

func (this StudentTicket) AddStudentTicket(ticket *db.StudentTicket) error {
	var err error
	var existingTicket db.StudentTicket
	this.db.First(&existingTicket, "serial_number = ?", ticket.SerialNumber)
	if existingTicket.Id != 0 {
		fmt.Println("Ticket with this serial number already exists")
		err = fmt.Errorf("Ticket with this serial number already exists")
	} else {
		this.db.Create(&ticket)
	}
	return err
}

func (this StudentTicket) GetStudentTickets() []db.StudentTicket {
	var tickets []db.StudentTicket
	this.db.Find(&tickets)
	return tickets
}

func (this StudentTicket) GetStudentTicketBySerialNumber(serialNumber string) db.StudentTicket {
	var ticket db.StudentTicket
	this.db.First(&ticket, "serial_number = ?", serialNumber)
	return ticket
}

func (this StudentTicket) GetStudentTicketById(id uint) db.StudentTicket {
	var ticket db.StudentTicket
	this.db.First(&ticket, "id = ?", id)
	return ticket
}
