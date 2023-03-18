package logic

import (
	"fmt"

	"github.com/it-02/dormitory/db"
	"github.com/it-02/dormitory/models"
)

func AddStudentTicket(ticket *models.StudentTicket) error {
	var err error
	var existingTicket models.StudentTicket
	db.DB.First(&existingTicket, "serial_number = ?", ticket.SerialNumber)
	if existingTicket.Id != 0 {
		fmt.Println("Ticket with this serial number already exists")
		err = fmt.Errorf("Ticket with this serial number already exists")
	} else {
		db.DB.Create(&ticket)
	}
	return err
}

func GetStudentTickets() []models.StudentTicket {
	var tickets []models.StudentTicket
	db.DB.Find(&tickets)
	return tickets
}

func GetStudentTicketBySerialNumber(serialNumber string) models.StudentTicket {
	var ticket models.StudentTicket
	db.DB.First(&ticket, "serial_number = ?", serialNumber)
	return ticket
}
