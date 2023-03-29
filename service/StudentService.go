package service

import (
	"fmt"
	"time"

	"github.com/it-02/dormitory/db"
	"github.com/it-02/dormitory/repository"
)

func RegisterStudent(student *db.Student, student_ticket *db.StudentTicket) error {
	ret := repository.AddStudentTicket(student_ticket)
	if ret != nil {
		return ret
	}

	student.StudentTicketId = student_ticket.Id
	ret = repository.AddStudent(student)
	if ret != nil {
		return ret
	}

	return ret
}

func SignContract(student_id uint, student_ticket_number string) error {
	var ret error
	ret = nil

	// if contract is already signed
	student := repository.GetStudentById(student_id)
	if student.ContractId != 0 {
		var contract db.Contract
		err := repository.GetContractById(student.ContractId, &contract)
		if err != nil {
			ret = err
			return ret
		} else {
			// so the contract exists
			// remove the contract from the student
			student.ContractId = 0
			err = repository.RemoveContractById(contract.Id)
			if err != nil {
				ret = err
				return ret
			}
		}
	}
	
	new_contract := repository.AddContract()

	// set the contract to the student
	err := repository.SetContract(student_id, new_contract.Id)
	if err != nil {
		ret = err
		return ret
	}

	return ret
}

func Settle(student_id uint, place_id uint) error {
	student := repository.GetStudentById(student_id)
	if student.PlaceId != 0 {
		return fmt.Errorf("Student is already settled, call resettle")
	}

	contract := db.Contract{}
	ret := repository.GetContractById(student.ContractId, &contract)
	if ret != nil {
		return ret
	}

	if contract.ExpireDate.Before(time.Now()) {
		ret = fmt.Errorf("Contract is expired")
		return ret
	}


	place := db.Place{}
	ret = repository.GetPlaceById(place_id, &place)
	if ret != nil {
		return ret
	}

	if !place.IsFree {
		ret = fmt.Errorf("Place is not free")
		return ret
	}

	// check room gender
	room := db.Room{}
	ret = repository.GetRoomById(place.RoomId, &room)
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
	student := repository.GetStudentById(student_id)
	if student.PlaceId == 0 {
		return fmt.Errorf("Student is not settled")
	}

	place := db.Place{}
	ret := repository.GetPlaceById(student.PlaceId, &place)
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
	place := db.Place{}
	ret := repository.GetPlaceById(place_id, &place)
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


