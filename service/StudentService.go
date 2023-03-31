package service

import (
	"github.com/it-02/dormitory/db"
	"github.com/it-02/dormitory/repository"
)

func RegisterStudent(student *db.Student, student_ticket *db.StudentTicket) error {
	ret := repository.AddStudentTicket(student_ticket)
	if ret != nil {
		return ret
	}

	student.StudentTicketId = student_ticket.Id
	ret = repository.AddStudent(student)
	if ret != nil {
		return ret
	}

	return ret
}

func SignContract(student_ticket_number string) error {
	return nil
}

func Settle(student_ticket_number string, roomNumber string) error {
	return nil
}

func Unsettle(student_ticket_number string) error {
	return nil
}

func Resettle(student_ticket_number string, roomNumber string) error {
	return nil
}
