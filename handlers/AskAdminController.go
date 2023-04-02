package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/it-02/dormitory/service"
)

type AskAdminRegisterRequest struct {
	Name    string
	Surname string
	IsMale  bool
	StudentTicketNumber string
	StudentTicketExpireDate time.Time
}

func AskAdminRegisterHandler(w http.ResponseWriter, r *http.Request) {
	var request AskAdminRegisterRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = service.AskAdminRegister(request.Name, request.Surname, request.IsMale, request.StudentTicketNumber, request.StudentTicketExpireDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

type AskAdminSignContractRequest struct {
	StudentTicketNumber string
}

func AskAdminSignContractHandler(w http.ResponseWriter, r *http.Request) {
	var request AskAdminSignContractRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = service.AskAdminSignContract(request.StudentTicketNumber)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

type AskAdminUnsettleRequest struct {
	StudentTicketNumber string
}

func AskAdminUnsettleHandler(w http.ResponseWriter, r *http.Request) {
	var request AskAdminUnsettleRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = service.AskAdminUnsettle(request.StudentTicketNumber)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

type AskAdminSettleRequest struct {
	StudentTicketNumber string
	RoomNumber string
}

func AskAdminSettleHandler(w http.ResponseWriter, r *http.Request) {
	var request AskAdminSettleRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = service.AskAdminSettle(request.StudentTicketNumber, request.RoomNumber)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

type AskAdminResettleRequest struct {
	StudentTicketNumber string
	RoomNumber string
}

func AskAdminResettleHandler(w http.ResponseWriter, r *http.Request) {
	var request AskAdminResettleRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = service.AskAdminResettle(request.StudentTicketNumber, request.RoomNumber)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

type AskAdminGetActionsRequest struct {
	UUID string `json:"uuid"`
}

func GetActionsHandler(w http.ResponseWriter, r *http.Request) {
	var request AskAdminGetActionsRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	actions, err := service.GetActions(request.UUID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(actions)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

type ResolveActionRequest struct {
	UUID string `json:"uuid"`
	ActionId uint `json:"action_id"`
	IsApproved bool `json:"is_approved"`
}

func ResolveActionHandler(w http.ResponseWriter, r *http.Request) {
	var request ResolveActionRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = service.ResolveAction(request.UUID, request.ActionId, request.IsApproved)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
