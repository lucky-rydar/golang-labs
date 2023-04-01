package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/it-02/dormitory/service"
)

type GetStudentsRequest struct {
	UUID string
}

func GetStudentsHandler(w http.ResponseWriter, r *http.Request) {
	var request GetStudentsRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err, students := service.GetStudents(request.UUID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(students)
}
