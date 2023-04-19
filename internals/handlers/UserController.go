package handlers

import (
	"encoding/json"
	"net/http"
)

type IUserService interface {
	RegisterUser(name string, pass string) error
	LoginUser(name string, pass string) (string, error)
	IsUserAdmin(uuid string) bool
}

type UserController struct {
	user_service IUserService
}

func NewUserController(user_service IUserService) *UserController {
	return &UserController{
		user_service: user_service,
	}
}

type RegisterUserRequest struct {
	Username string
	Password string
}

func (uc *UserController) RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	var request RegisterUserRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = uc.user_service.RegisterUser(request.Username, request.Password)
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

func (uc *UserController) LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	var request LoginUserRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user_uuid, err := uc.user_service.LoginUser(request.Username, request.Password)
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

