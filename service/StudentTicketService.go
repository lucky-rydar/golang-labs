package service

import (
	"github.com/it-02/dormitory/repository"
	"github.com/it-02/dormitory/db"
)

type IStudentTicketService interface {
	AddStudentTicket(studentTicket *db.StudentTicket) error
	GetStudentTickets() []db.StudentTicket
	GetStudentTicketBySerialNumber(serialNumber string) db.StudentTicket
}

type StudentTicketService struct {
	student_ticket_repository repository.IStudentTicket
}

func NewStudentTicketService(student_ticket_repository repository.IStudentTicket) *StudentTicketService {
	return &StudentTicketService{
		student_ticket_repository: student_ticket_repository,
	}
}

func (this StudentTicketService) AddStudentTicket(studentTicket *db.StudentTicket) error {
	return this.student_ticket_repository.AddStudentTicket(studentTicket)
}

func (this StudentTicketService) GetStudentTickets() []db.StudentTicket {
	return this.student_ticket_repository.GetStudentTickets()
}

func (this StudentTicketService) GetStudentTicketBySerialNumber(serialNumber string) db.StudentTicket {
	return this.student_ticket_repository.GetStudentTicketBySerialNumber(serialNumber)
}


