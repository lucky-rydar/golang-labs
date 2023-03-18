package logic

import (
	"fmt"
	"time"

	"github.com/it-02/dormitory/db"
	"github.com/it-02/dormitory/models"
)

func AddStudent(student *models.Student) error {
	var err error
	var existingStudent models.Student
	db.DB.First(&existingStudent, "name = ? AND surname = ?", student.Name, student.Surname)
	if existingStudent.Id != 0 {
		fmt.Println("Student with this name and surname already exists")
		err = fmt.Errorf("Student with this name and surname already exists")
	} else {
		db.DB.Create(&student)
	}
	return err
}

func RegisterStudent(student *models.Student, student_ticket *models.StudentTicket) error {
	var ret error
	ret = nil

	err := AddStudentTicket(student_ticket)
	if err != nil {
		ret = err
		return ret
	}

	student.StudentTicketId = student_ticket.Id
	err = AddStudent(student)
	if err != nil {
		ret = err
		return ret
	}

	return ret
}

func SetContract(student_id uint, contract_id uint) error {
	var ret error
	ret = nil

	student := GetStudentById(student_id)
	student.ContractId = contract_id
	db.DB.Save(&student)

	return ret
}

func SignContract(student_id uint, student_ticket_number string) error {
	var ret error
	ret = nil

	// if contract is already signed
	student := GetStudentById(student_id)
	if student.ContractId != 0 {
		var contract models.Contract
		err := GetContractById(student.ContractId, &contract)
		if err != nil {
			ret = err
			return ret
		} else {
			// so the contract exists
			// remove the contract from the student
			student.ContractId = 0
			err = RemoveContractById(contract.Id)
			if err != nil {
				ret = err
				return ret
			}
		}
	}
	
	new_contract := AddContract()

	// set the contract to the student
	err := SetContract(student_id, new_contract.Id)
	if err != nil {
		ret = err
		return ret
	}

	return ret
}

func Settle(student_id uint, place_id uint) error {
	var ret error

	student := GetStudentById(student_id)
	if student.PlaceId != 0 {
		return fmt.Errorf("Student is already settled, call resettle")
	}

	contract := models.Contract{}
	err := GetContractById(student.ContractId, &contract)
	if err != nil {
		ret = err
		return ret
	}

	if contract.ExpireDate.Before(time.Now()) {
		ret = fmt.Errorf("Contract is expired")
		return ret
	}


	place := models.Place{}
	err = GetPlaceById(place_id, &place)
	if err != nil {
		ret = err
		return ret
	}

	if !place.IsFree {
		ret = fmt.Errorf("Place is not free")
		return ret
	}

	student.PlaceId = place_id
	place.IsFree = false

	db.DB.Save(&student)
	db.DB.Save(&place)

	return ret
}

func Unsettle(student_id uint) error {
	var ret error

	student := GetStudentById(student_id)
	if student.PlaceId == 0 {
		return fmt.Errorf("Student is not settled")
	}

	place := models.Place{}
	err := GetPlaceById(student.PlaceId, &place)
	if err != nil {
		ret = err
		return ret
	}

	student.PlaceId = 0
	place.IsFree = true

	db.DB.Save(&student)
	db.DB.Save(&place)

	return ret
}

func Resettle(student_id uint, place_id uint) error {
	var ret error

	// if place is not free return error
	place := models.Place{}
	err := GetPlaceById(place_id, &place)
	if err != nil {
		ret = err
		return ret
	}

	if !place.IsFree {
		ret = fmt.Errorf("Place is not free")
		return ret
	}

	// unsettle the student
	err = Unsettle(student_id)
	if err != nil {
		ret = err
		return ret
	}

	// settle the student
	err = Settle(student_id, place_id)
	if err != nil {
		ret = err
		return ret
	}

	return ret
}

func GetStudents() []models.Student {
	var students []models.Student
	db.DB.Find(&students)
	return students
}

func GetStudentById(id uint) models.Student {
	var student models.Student
	db.DB.Where("id = ?", id).First(&student)
	return student
}
