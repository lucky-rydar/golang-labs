package handlers

import (
	"net/http"
	"encoding/json"

	"github.com/it-02/dormitory/service"
)

func GetDormitoryLoadHandler(w http.ResponseWriter, r *http.Request) {
	dormitoryLoad := service.GetDormitoryLoad()
	err := json.NewEncoder(w).Encode(dormitoryLoad)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
