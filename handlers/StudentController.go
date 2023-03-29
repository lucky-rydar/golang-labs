package handlers

import (
	"net/http"
	"encoding/json"
	"time"

	"github.com/it-02/dormitory/repository"
	"github.com/it-02/dormitory/service"
	"github.com/it-02/dormitory/db"
)

type AddSrudentRequest struct {
	Name    string
	Surname string
	IsMale  bool
	StudentTicketNumber string
	StudentTicketExpireDate time.Time
}

func RegisterStudentHandler(w http.ResponseWriter, r *http.Request) {
	var request AddSrudentRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	student := db.Student{
		Name:    request.Name,
		Surname: request.Surname,
		IsMale:  request.IsMale,
		ContractId: 0,
		StudentTicketId: 0,
		PlaceId: 0,
	}

	student_ticket := db.StudentTicket{
		SerialNumber: request.StudentTicketNumber,
		ExpireDate:   request.StudentTicketExpireDate,
	}

	err = service.RegisterStudent(&student, &student_ticket)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(student)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

type SignContractRequest struct {
	StudentId uint
	StudentTicketNumber string
}

func SignContractHandler(w http.ResponseWriter, r *http.Request) {
	var request SignContractRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = service.SignContract(request.StudentId, request.StudentTicketNumber)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

type SettleRequest struct {
	StudentId uint
	PlaceId uint
}

func SettleHandler(w http.ResponseWriter, r *http.Request) {
	var request SettleRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = service.Settle(request.StudentId, request.PlaceId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

type UnsettleRequest struct {
	StudentId uint
}

func UnsettleHandler(w http.ResponseWriter, r *http.Request) {
	var request UnsettleRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = service.Unsettle(request.StudentId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

type ResettleRequest struct {
	StudentId uint
	PlaceId uint
}

func ResettleHandler(w http.ResponseWriter, r *http.Request) {
	var request ResettleRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = service.Resettle(request.StudentId, request.PlaceId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func GetStudentsHandler(w http.ResponseWriter, r *http.Request) {
	students := repository.GetStudents()
	err := json.NewEncoder(w).Encode(students)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
