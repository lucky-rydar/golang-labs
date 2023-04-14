package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/it-02/dormitory/service"
)

type IAskAdminController interface {
	AskAdminRegisterHandler(w http.ResponseWriter, r *http.Request)
	AskAdminSignContractHandler(w http.ResponseWriter, r *http.Request)
	AskAdminUnsettleHandler(w http.ResponseWriter, r *http.Request)
	AskAdminSettleHandler(w http.ResponseWriter, r *http.Request)
	AskAdminResettleHandler(w http.ResponseWriter, r *http.Request)
	GetActionsHandler(w http.ResponseWriter, r *http.Request)
	ResolveActionHandler(w http.ResponseWriter, r *http.Request)
}

type AskAdminController struct {
	ask_admin_service *this.ask_admin_service.IAskAdminService
}

func NewAskAdminController(ask_admin_service *this.ask_admin_service.IAskAdminService) IAskAdminController {
	return &AskAdminController{ask_admin_service: ask_admin_service}
}

type AskAdminRegisterRequest struct {
	Name    string
	Surname string
	IsMale  bool
	StudentTicketNumber string
	StudentTicketExpireDate time.Time
}

func (this AskAdminController) AskAdminRegisterHandler(w http.ResponseWriter, r *http.Request) {
	var request AskAdminRegisterRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = this.ask_admin_service.AskAdminRegister(request.Name, request.Surname, request.IsMale, request.StudentTicketNumber, request.StudentTicketExpireDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

type AskAdminSignContractRequest struct {
	StudentTicketNumber string
}

func (this AskAdminController) AskAdminSignContractHandler(w http.ResponseWriter, r *http.Request) {
	var request AskAdminSignContractRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = this.ask_admin_service.AskAdminSignContract(request.StudentTicketNumber)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

type AskAdminUnsettleRequest struct {
	StudentTicketNumber string
}

func (this AskAdminController) AskAdminUnsettleHandler(w http.ResponseWriter, r *http.Request) {
	var request AskAdminUnsettleRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = this.ask_admin_service.AskAdminUnsettle(request.StudentTicketNumber)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

type AskAdminSettleRequest struct {
	StudentTicketNumber string
	RoomNumber string
}

func (this AskAdminController) AskAdminSettleHandler(w http.ResponseWriter, r *http.Request) {
	var request AskAdminSettleRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = this.ask_admin_service.AskAdminSettle(request.StudentTicketNumber, request.RoomNumber)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

type AskAdminResettleRequest struct {
	StudentTicketNumber string
	RoomNumber string
}

func (this AskAdminController) AskAdminResettleHandler(w http.ResponseWriter, r *http.Request) {
	var request AskAdminResettleRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = this.ask_admin_service.AskAdminResettle(request.StudentTicketNumber, request.RoomNumber)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

type AskAdminGetActionsRequest struct {
	UUID string `json:"uuid"`
}

func (this AskAdminController) GetActionsHandler(w http.ResponseWriter, r *http.Request) {
	var request AskAdminGetActionsRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	actions, err := this.ask_admin_service.GetActions(request.UUID)
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

func (this AskAdminController) ResolveActionHandler(w http.ResponseWriter, r *http.Request) {
	var request ResolveActionRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = this.ask_admin_service.ResolveAction(request.UUID, request.ActionId, request.IsApproved)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
