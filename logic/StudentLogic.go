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
	ret := AddStudentTicket(student_ticket)
	if ret != nil {
		return ret
	}

	student.StudentTicketId = student_ticket.Id
	ret = AddStudent(student)
	if ret != nil {
		return ret
	}

	return ret
}

func SetContract(student_id uint, contract_id uint) error {
	student := GetStudentById(student_id)
	student.ContractId = contract_id
	db.DB.Save(&student)

	return nil
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
	student := GetStudentById(student_id)
	if student.PlaceId != 0 {
		return fmt.Errorf("Student is already settled, call resettle")
	}

	contract := models.Contract{}
	ret := GetContractById(student.ContractId, &contract)
	if ret != nil {
		return ret
	}

	if contract.ExpireDate.Before(time.Now()) {
		ret = fmt.Errorf("Contract is expired")
		return ret
	}


	place := models.Place{}
	ret = GetPlaceById(place_id, &place)
	if ret != nil {
		return ret
	}

	if !place.IsFree {
		ret = fmt.Errorf("Place is not free")
		return ret
	}

	// check room gender
	room := models.Room{}
	ret = GetRoomById(place.RoomId, &room)
	if ret != nil {
		return ret
	}

	if room.IsMale != student.IsMale {
		ret = fmt.Errorf("Student is not the same gender as the room")
		return ret
	}

	student.PlaceId = place_id
	place.IsFree = false

	db.DB.Save(&student)
	db.DB.Save(&place)

	return ret
}

func Unsettle(student_id uint) error {
	student := GetStudentById(student_id)
	if student.PlaceId == 0 {
		return fmt.Errorf("Student is not settled")
	}

	place := models.Place{}
	ret := GetPlaceById(student.PlaceId, &place)
	if ret != nil {
		return ret
	}

	student.PlaceId = 0
	place.IsFree = true

	db.DB.Save(&student)
	db.DB.Save(&place)

	return ret
}

func Resettle(student_id uint, place_id uint) error {
	// if place is not free return error
	place := models.Place{}
	ret := GetPlaceById(place_id, &place)
	if ret != nil {
		return ret
	}

	if !place.IsFree {
		ret = fmt.Errorf("Place is not free")
		return ret
	}

	// unsettle the student
	ret = Unsettle(student_id)
	if ret != nil {
		return ret
	}

	// settle the student
	ret = Settle(student_id, place_id)
	if ret != nil {
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
