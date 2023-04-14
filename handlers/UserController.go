package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/it-02/dormitory/service"
)

type IUserController interface {
	RegisterUserHandler(w http.ResponseWriter, r *http.Request)
	LoginUserHandler(w http.ResponseWriter, r *http.Request)
}

type UserController struct {
	user_service service.IUserService
}

func NewUserController(user_service service.IUserService) *UserController {
	return &UserController{
		user_service: user_service,
	}
}

type RegisterUserRequest struct {
	Username string
	Password string
}

func (this UserController) RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	var request RegisterUserRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = this.user_service.RegisterUser(request.Username, request.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

type LoginUserRequest struct {
	Username string
	Password string
}

type LoginResponse struct {
	Uuid string
}

func (this UserController) LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	var request LoginUserRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user_uuid, err := this.user_service.LoginUser(request.Username, request.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// encode response
	response := LoginResponse{
		Uuid: user_uuid,
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

