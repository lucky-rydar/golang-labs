package repository

import (
	"fmt"

	"github.com/it-02/dormitory/internals/db"
	"gorm.io/gorm"
)

type StudentTicket struct {
	db *gorm.DB
}

func NewStudentTicket(db *gorm.DB) *StudentTicket {
	return &StudentTicket{db: db}
}

func (st *StudentTicket) AddStudentTicket(ticket *db.StudentTicket) error {
	var err error
	var existingTicket db.StudentTicket
	st.db.First(&existingTicket, "serial_number = ?", ticket.SerialNumber)
	if existingTicket.Id != 0 {
		fmt.Println("Ticket with st serial number already exists")
		err = fmt.Errorf("Ticket with st serial number already exists")
	} else {
		st.db.Create(&ticket)
	}
	return err
}

func (st *StudentTicket) GetStudentTickets() []db.StudentTicket {
	var tickets []db.StudentTicket
	st.db.Find(&tickets)
	return tickets
}

func (st *StudentTicket) GetStudentTicketBySerialNumber(serialNumber string) db.StudentTicket {
	var ticket db.StudentTicket
	st.db.First(&ticket, "serial_number = ?", serialNumber)
	return ticket
}

func (st *StudentTicket) GetStudentTicketById(id uint) db.StudentTicket {
	var ticket db.StudentTicket
	st.db.First(&ticket, "id = ?", id)
	return ticket
}
