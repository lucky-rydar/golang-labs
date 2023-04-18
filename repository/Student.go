package repository

import (
	"fmt"

	"github.com/it-02/dormitory/db"
	"gorm.io/gorm"
)

type IStudent interface {
	AddStudent(student *db.Student) error
	SetContract(student_id uint, contract_id uint) error
	GetStudents() []db.Student
	GetStudentById(id uint) db.Student
	GetStudentByTicketId(ticket_id uint) db.Student
	SetStudentToPlace(student_id uint, place_id uint) error
	UnsetStudentFromPlace(student_id uint) error
	GetStudentsByPlaceIds(place_ids []uint) []db.Student
}

type Student struct {
	db *gorm.DB
	place_repository IPlace
}

func NewStudent(db *gorm.DB, place_repository IPlace) IStudent {
	return &Student{db: db, place_repository: place_repository}
}

func (s *Student) AddStudent(student *db.Student) error {
	var err error
	var existingStudent db.Student
	s.db.First(&existingStudent, "name = ? AND surname = ?", student.Name, student.Surname)
	if existingStudent.Id != 0 {
		fmt.Println("Student with s name and surname already exists")
		err = fmt.Errorf("Student with s name and surname already exists")
	} else {
		s.db.Create(&student)
	}
	return err
}

func (s *Student) SetContract(student_id uint, contract_id uint) error {
	student := s.GetStudentById(student_id)
	student.ContractId = contract_id
	s.db.Save(&student)

	return nil
}

func (s *Student) GetStudents() []db.Student {
	var students []db.Student
	s.db.Find(&students)
	return students
}

func (s *Student) GetStudentById(id uint) db.Student {
	var student db.Student
	s.db.Where("id = ?", id).First(&student)
	return student
}

func (s *Student) GetStudentByTicketId(ticket_id uint) db.Student {
	var student db.Student
	s.db.Where("student_ticket_id = ?", ticket_id).First(&student)
	return student
}

func (s *Student) SetStudentToPlace(student_id uint, place_id uint) error {
	student := s.GetStudentById(student_id)

	place := db.Place{}
	err := s.place_repository.GetPlaceById(place_id, &place)
	if err != nil {
		return err
	}

	student.PlaceId = place_id
	s.db.Save(&student)

	place.IsFree = false
	s.db.Save(&place)

	return nil
}

func (s *Student) UnsetStudentFromPlace(student_id uint) error {
	student := s.GetStudentById(student_id)
	
	// make place free
	place := db.Place{}
	err := s.place_repository.GetPlaceById(student.PlaceId, &place)
	if err != nil {
		return err
	}

	place.IsFree = true
	s.db.Save(&place)
	
	student.PlaceId = 0
	s.db.Save(&student)

	return nil
}

func (s *Student) GetStudentsByPlaceIds(place_ids []uint) []db.Student {
	var students []db.Student
	s.db.Where("place_id IN (?)", place_ids).Find(&students)
	return students
}
