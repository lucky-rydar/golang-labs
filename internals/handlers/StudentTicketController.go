package handlers

import (
	"net/http"
	"encoding/json"
	"time"

	"github.com/it-02/dormitory/internals/db"
)

type IStudentTicketService interface {
	AddStudentTicket(studentTicket *db.StudentTicket) error
	GetStudentTickets() []db.StudentTicket
	GetStudentTicketBySerialNumber(serialNumber string) db.StudentTicket
}

type StudentTicketController struct {
	student_ticket_service IStudentTicketService
}

func NewStudentTicketController(student_ticket_service IStudentTicketService) *StudentTicketController {
	return &StudentTicketController{
		student_ticket_service: student_ticket_service,
	}
}

type AddStudentTicketRequest struct {
	SerialNumber string
	ExpireDate   time.Time
}

func (stc *StudentTicketController) AddStudentTicketHandler(w http.ResponseWriter, r *http.Request) {
	var request AddStudentTicketRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ticket := db.StudentTicket{
		SerialNumber: request.SerialNumber,
		ExpireDate:   request.ExpireDate,
	}

	err = stc.student_ticket_service.AddStudentTicket(&ticket)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (stc *StudentTicketController) GetStudentTicketsHandler(w http.ResponseWriter, r *http.Request) {
	tickets := stc.student_ticket_service.GetStudentTickets()
	err := json.NewEncoder(w).Encode(tickets)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

type GetStudentTicketBySerialNumberRequest struct {
	SerialNumber string
}

func (stc *StudentTicketController) GetStudentTicketBySerialNumberHandler(w http.ResponseWriter, r *http.Request) {
	var request GetStudentTicketBySerialNumberRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ticket := stc.student_ticket_service.GetStudentTicketBySerialNumber(request.SerialNumber)
	err = json.NewEncoder(w).Encode(ticket)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}