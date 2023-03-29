package repository

import (
	"fmt"

	"github.com/it-02/dormitory/db"
)

func AddStudent(student *db.Student) error {
	var err error
	var existingStudent db.Student
	db.DB.First(&existingStudent, "name = ? AND surname = ?", student.Name, student.Surname)
	if existingStudent.Id != 0 {
		fmt.Println("Student with this name and surname already exists")
		err = fmt.Errorf("Student with this name and surname already exists")
	} else {
		db.DB.Create(&student)
	}
	return err
}

func SetContract(student_id uint, contract_id uint) error {
	student := GetStudentById(student_id)
	student.ContractId = contract_id
	db.DB.Save(&student)

	return nil
}

func GetStudents() []db.Student {
	var students []db.Student
	db.DB.Find(&students)
	return students
}

func GetStudentById(id uint) db.Student {
	var student db.Student
	db.DB.Where("id = ?", id).First(&student)
	return student
}
