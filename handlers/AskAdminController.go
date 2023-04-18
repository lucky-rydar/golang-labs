package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/it-02/dormitory/db"
)

type IAskAdminService interface {
	AskAdminRegister(name string, surname string, isMale bool, studentTicketNumber string, studentTicketExpireDate time.Time) error
	AskAdminSignContract(studentTicketNumber string) error
	AskAdminUnsettle(studentTicketNumber string) error
	AskAdminSettle(studentTicketNumber string, roomNumber string) error
	AskAdminResettle(studentTicketNumber string, roomNumber string) error
	GetActions(uuid string) ([]db.AskAdmin, error)
	ResolveAction(uuid string, actionId uint, isApproved bool) error
}

type AskAdminController struct {
	ask_admin_service IAskAdminService
}

func NewAskAdminController(ask_admin_service IAskAdminService) *AskAdminController {
	return &AskAdminController{ask_admin_service: ask_admin_service}
}

type AskAdminRegisterRequest struct {
	Name    string
	Surname string
	IsMale  bool
	StudentTicketNumber string
	StudentTicketExpireDate time.Time
}

func (aac *AskAdminController) AskAdminRegisterHandler(w http.ResponseWriter, r *http.Request) {
	var request AskAdminRegisterRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = aac.ask_admin_service.AskAdminRegister(request.Name, request.Surname, request.IsMale, request.StudentTicketNumber, request.StudentTicketExpireDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

type AskAdminSignContractRequest struct {
	StudentTicketNumber string
}

func (aac *AskAdminController) AskAdminSignContractHandler(w http.ResponseWriter, r *http.Request) {
	var request AskAdminSignContractRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = aac.ask_admin_service.AskAdminSignContract(request.StudentTicketNumber)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

type AskAdminUnsettleRequest struct {
	StudentTicketNumber string
}

func (aac *AskAdminController) AskAdminUnsettleHandler(w http.ResponseWriter, r *http.Request) {
	var request AskAdminUnsettleRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = aac.ask_admin_service.AskAdminUnsettle(request.StudentTicketNumber)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

type AskAdminSettleRequest struct {
	StudentTicketNumber string
	RoomNumber string
}

func (aac *AskAdminController) AskAdminSettleHandler(w http.ResponseWriter, r *http.Request) {
	var request AskAdminSettleRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = aac.ask_admin_service.AskAdminSettle(request.StudentTicketNumber, request.RoomNumber)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

type AskAdminResettleRequest struct {
	StudentTicketNumber string
	RoomNumber string
}

func (aac *AskAdminController) AskAdminResettleHandler(w http.ResponseWriter, r *http.Request) {
	var request AskAdminResettleRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = aac.ask_admin_service.AskAdminResettle(request.StudentTicketNumber, request.RoomNumber)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

type AskAdminGetActionsRequest struct {
	UUID string `json:"uuid"`
}

func (aac *AskAdminController) GetActionsHandler(w http.ResponseWriter, r *http.Request) {
	var request AskAdminGetActionsRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	actions, err := aac.ask_admin_service.GetActions(request.UUID)
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

func (aac *AskAdminController) ResolveActionHandler(w http.ResponseWriter, r *http.Request) {
	var request ResolveActionRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = aac.ask_admin_service.ResolveAction(request.UUID, request.ActionId, request.IsApproved)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
