package repository

import (
	"fmt"

	"github.com/it-02/dormitory/db"
)

type IStudent interface {
	AddStudent(student *db.Student) error
	SetContract(student_id uint, contract_id uint) error
	GetStudents() []db.Student
	GetStudentById(id uint) db.Student
	GetStudentByTicketId(ticket_id uint) db.Student
	SetStudentToPlace(student_id uint, place_id uint) error
	UnsetStudentFromPlace(student_id uint) error
}

type Student struct {
	db *gorm.DB
}

func NewStudent(db *gorm.DB) IStudent {
	return &Student{db: db}
}

func (this Student) AddStudent(student *db.Student) error {
	var err error
	var existingStudent db.Student
	this.db.First(&existingStudent, "name = ? AND surname = ?", student.Name, student.Surname)
	if existingStudent.Id != 0 {
		fmt.Println("Student with this name and surname already exists")
		err = fmt.Errorf("Student with this name and surname already exists")
	} else {
		this.db.Create(&student)
	}
	return err
}

func (this Student) SetContract(student_id uint, contract_id uint) error {
	student := GetStudentById(student_id)
	student.ContractId = contract_id
	this.db.Save(&student)

	return nil
}

func (this Student) GetStudents() []db.Student {
	var students []db.Student
	this.db.Find(&students)
	return students
}

func (this Student) GetStudentById(id uint) db.Student {
	var student db.Student
	this.db.Where("id = ?", id).First(&student)
	return student
}

func (this Student) GetStudentByTicketId(ticket_id uint) db.Student {
	var student db.Student
	this.db.Where("student_ticket_id = ?", ticket_id).First(&student)
	return student
}

func (this Student) SetStudentToPlace(student_id uint, place_id uint) error {
	student := GetStudentById(student_id)

	place := db.Place{}
	err := GetPlaceById(place_id, &place)
	if err != nil {
		return err
	}

	student.PlaceId = place_id
	this.db.Save(&student)

	place.IsFree = false
	this.db.Save(&place)

	return nil
}

func (this Student) UnsetStudentFromPlace(student_id uint) error {
	student := GetStudentById(student_id)
	
	// make place free
	place := db.Place{}
	err := GetPlaceById(student.PlaceId, &place)
	if err != nil {
		return err
	}

	place.IsFree = true
	this.db.Save(&place)
	
	student.PlaceId = 0
	this.db.Save(&student)

	return nil
}
