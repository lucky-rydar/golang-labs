package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/it-02/dormitory/service"
)

type RegisterUserRequest struct {
	Username string
	Password string
}

func RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	var request RegisterUserRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = service.RegisterUser(request.Username, request.Password)
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

func LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	var request LoginUserRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user_uuid, err := service.LoginUser(request.Username, request.Password)
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

