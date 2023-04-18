package handlers

import (
	"net/http"
	"encoding/json"

	"github.com/it-02/dormitory/service"
)

type DormitoryLoadController struct {
	dormitory_load_service service.IDormitoryLoadService
}

func NewDormitoryLoadController(dormitory_load_service service.IDormitoryLoadService) *DormitoryLoadController {
	return &DormitoryLoadController{
		dormitory_load_service: dormitory_load_service,
	}
}

type GetDormStatsRequest struct {
	UUID string `json:"uuid"`
}

func (dlc *DormitoryLoadController) GetDormitoryLoadHandler(w http.ResponseWriter, r *http.Request) {
	var request GetDormStatsRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	dormitoryLoad, err := dlc.dormitory_load_service.GetDormitoryLoad(request.UUID)
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
