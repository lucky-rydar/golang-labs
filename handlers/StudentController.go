package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/it-02/dormitory/service"
)

type IStudentController interface {
	GetStudentsHandler(w http.ResponseWriter, r *http.Request)
}

type StudentController struct {
	student_service service.IStudentService
}

func NewStudentController(student_service service.IStudentService) *StudentController {
	return &StudentController{
		student_service: student_service,
	}
}

type GetStudentsRequest struct {
	UUID string
}

func (this StudentController) GetStudentsHandler(w http.ResponseWriter, r *http.Request) {
	var request GetStudentsRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err, students := this.student_service.GetStudents(request.UUID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(students)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
