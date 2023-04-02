package handlers

import (
	"net/http"
	"encoding/json"

	"github.com/it-02/dormitory/service"
)

type GetDormStatsRequest struct {
	UUID string `json:"uuid"`
}

func GetDormitoryLoadHandler(w http.ResponseWriter, r *http.Request) {
	var request GetDormStatsRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	dormitoryLoad, err := service.GetDormitoryLoad(request.UUID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(dormitoryLoad)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
