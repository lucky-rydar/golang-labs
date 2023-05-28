package service

import (
	"github.com/it-02/dormitory/internals/db"
)

type IStudentTicket interface {
	AddStudentTicket(ticket *db.StudentTicket) error
	GetStudentTickets() []db.StudentTicket
	GetStudentTicketBySerialNumber(serialNumber string) db.StudentTicket
}

type StudentTicketService struct {
	student_ticket_repository IStudentTicket
}

func NewStudentTicketService(student_ticket_repository IStudentTicket) *StudentTicketService {
	return &StudentTicketService{
		student_ticket_repository: student_ticket_repository,
	}
}

func (sts *StudentTicketService) AddStudentTicket(studentTicket *db.StudentTicket) error {
	return sts.student_ticket_repository.AddStudentTicket(studentTicket)
}

func (sts *StudentTicketService) GetStudentTickets() []db.StudentTicket {
	return sts.student_ticket_repository.GetStudentTickets()
}

func (sts *StudentTicketService) GetStudentTicketBySerialNumber(serialNumber string) db.StudentTicket {
	return sts.student_ticket_repository.GetStudentTicketBySerialNumber(serialNumber)
}


