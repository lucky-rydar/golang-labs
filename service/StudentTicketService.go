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
	student_ticket_repository *repository.IStudentTicket
}

func NewStudentTicketService() *StudentTicketService {
	return &StudentTicketService{
		student_ticket_repository: repository.NewStudentTicketRepository(),
	}
}

func (this StudentTicketService) AddStudentTicket(studentTicket *db.StudentTicket) error {
	return repository.AddStudentTicket(studentTicket)
}

func (this StudentTicketService) GetStudentTickets() []db.StudentTicket {
	return repository.GetStudentTickets()
}

func (this StudentTicketService) GetStudentTicketBySerialNumber(serialNumber string) db.StudentTicket {
	return repository.GetStudentTicketBySerialNumber(serialNumber)
}


