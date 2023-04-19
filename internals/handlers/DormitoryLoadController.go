package handlers

import (
	"net/http"
	"encoding/json"

	"github.com/it-02/dormitory/service"
)

type Istructs.DormitoryLoadService interface {
	Getstructs.DormitoryLoad(uuid string) (service.structs.DormitoryLoad, error)
}

type structs.DormitoryLoadController struct {
	dormitory_load_service Istructs.DormitoryLoadService
}

func Newstructs.DormitoryLoadController(dormitory_load_service Istructs.DormitoryLoadService) *structs.DormitoryLoadController {
	return &structs.DormitoryLoadController{
		dormitory_load_service: dormitory_load_service,
	}
}

type GetDormStatsRequest struct {
	UUID string `json:"uuid"`
}

func (dlc *structs.DormitoryLoadController) Getstructs.DormitoryLoadHandler(w http.ResponseWriter, r *http.Request) {
	var request GetDormStatsRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	dormitoryLoad, err := dlc.dormitory_load_service.Getstructs.DormitoryLoad(request.UUID)
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
