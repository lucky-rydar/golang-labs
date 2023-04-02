package repository

import (
	"fmt"

	"github.com/it-02/dormitory/db"
)

func AddStudentTicket(ticket *db.StudentTicket) error {
	var err error
	var existingTicket db.StudentTicket
	db.DB.First(&existingTicket, "serial_number = ?", ticket.SerialNumber)
	if existingTicket.Id != 0 {
		fmt.Println("Ticket with this serial number already exists")
		err = fmt.Errorf("Ticket with this serial number already exists")
	} else {
		db.DB.Create(&ticket)
	}
	return err
}

func GetStudentTickets() []db.StudentTicket {
	var tickets []db.StudentTicket
	db.DB.Find(&tickets)
	return tickets
}

func GetStudentTicketBySerialNumber(serialNumber string) db.StudentTicket {
	var ticket db.StudentTicket
	db.DB.First(&ticket, "serial_number = ?", serialNumber)
	return ticket
}

func GetStudentTicketById(id uint) db.StudentTicket {
	var ticket db.StudentTicket
	db.DB.First(&ticket, "id = ?", id)
	return ticket
}
