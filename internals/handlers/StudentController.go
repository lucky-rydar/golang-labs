package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/it-02/dormitory/internals/db"
	"github.com/it-02/dormitory/internals/features/student/structs"
)

type IStudentService interface {
	GetStudents(uuid string) (error, []structs.StudentRepr)
}

type StudentController struct {
	student_service IStudentService
}

func NewStudentController(student_service IStudentService) *StudentController {
	return &StudentController{
		student_service: student_service,
	}
}

type GetStudentsRequest struct {
	UUID string
}

func (sc *StudentController) GetStudentsHandler(w http.ResponseWriter, r *http.Request) {
	var request GetStudentsRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err, students := sc.student_service.GetStudents(request.UUID)
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
