package handlers

import (
	"net/http"
	"encoding/json"
	"time"

	"github.com/it-02/dormitory/repository"
	"github.com/it-02/dormitory/db"
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

	ticket := db.StudentTicket{
		SerialNumber: request.SerialNumber,
		ExpireDate:   request.ExpireDate,
	}

	err = repository.AddStudentTicket(&ticket)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func GetStudentTicketsHandler(w http.ResponseWriter, r *http.Request) {
	tickets := repository.GetStudentTickets()
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

	ticket := repository.GetStudentTicketBySerialNumber(request.SerialNumber)
	err = json.NewEncoder(w).Encode(ticket)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}