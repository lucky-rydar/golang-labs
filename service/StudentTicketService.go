package service

import (
	"github.com/it-02/dormitory/repository"
	"github.com/it-02/dormitory/db"
)

func AddStudentTicket(studentTicket *db.StudentTicket) error {
	return repository.AddStudentTicket(studentTicket)
}

func GetStudentTickets() []db.StudentTicket {
	return repository.GetStudentTickets()
}

func GetStudentTicketBySerialNumber(serialNumber string) db.StudentTicket {
	return repository.GetStudentTicketBySerialNumber(serialNumber)
}


