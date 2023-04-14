package handlers

import (
	"net/http"
	"encoding/json"
	"time"

	"github.com/it-02/dormitory/service"
	"github.com/it-02/dormitory/db"
)

type IStudentTicketController interface {
	AddStudentTicketHandler(w http.ResponseWriter, r *http.Request)
	GetStudentTicketsHandler(w http.ResponseWriter, r *http.Request)
	GetStudentTicketBySerialNumberHandler(w http.ResponseWriter, r *http.Request)
}

type StudentTicketController struct {
	student_ticket_service *service.IStudentTicketService
}

func NewStudentTicketController(student_ticket_service *service.IStudentTicketService) *StudentTicketController {
	return &StudentTicketController{
		student_ticket_service: student_ticket_service,
	}
}

type AddStudentTicketRequest struct {
	SerialNumber string
	ExpireDate   time.Time
}

func (this StudentTicketController) AddStudentTicketHandler(w http.ResponseWriter, r *http.Request) {
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

	err = this.student_ticket_service.AddStudentTicket(&ticket)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (this StudentTicketController) GetStudentTicketsHandler(w http.ResponseWriter, r *http.Request) {
	tickets := this.student_ticket_service.GetStudentTickets()
	err := json.NewEncoder(w).Encode(tickets)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

type GetStudentTicketBySerialNumberRequest struct {
	SerialNumber string
}

func (this StudentTicketController) GetStudentTicketBySerialNumberHandler(w http.ResponseWriter, r *http.Request) {
	var request GetStudentTicketBySerialNumberRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ticket := this.student_ticket_service.GetStudentTicketBySerialNumber(request.SerialNumber)
	err = json.NewEncoder(w).Encode(ticket)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}