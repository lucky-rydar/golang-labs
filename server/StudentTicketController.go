package server

import (
	"net/http"
	"encoding/json"
	"time"

	"github.com/it-02/dormitory/models"
	"github.com/it-02/dormitory/logic"
)

type AddStudentTicketRequest struct {
	SerialNumber string
	ExpireDate   time.Time
}

func AddStudentTicketHandler(w http.ResponseWriter, r *http.Request) {
	var request AddStudentTicketRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ticket := models.StudentTicket{
		SerialNumber: request.SerialNumber,
		ExpireDate:   request.ExpireDate,
	}

	err = logic.AddStudentTicket(ticket)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func GetStudentTicketsHandler(w http.ResponseWriter, r *http.Request) {
	tickets := logic.GetStudentTickets()
	err := json.NewEncoder(w).Encode(tickets)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

type GetStudentTicketBySerialNumberRequest struct {
	SerialNumber string
}

func GetStudentTicketBySerialNumberHandler(w http.ResponseWriter, r *http.Request) {
	var request GetStudentTicketBySerialNumberRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ticket := logic.GetStudentTicketBySerialNumber(request.SerialNumber)
	err = json.NewEncoder(w).Encode(ticket)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}