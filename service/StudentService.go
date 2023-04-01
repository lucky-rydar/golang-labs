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

func SignContract(student_ticket_number string) error {
	ticket := repository.GetStudentTicketBySerialNumber(student_ticket_number)
	if ticket.Id == 0 {
		return fmt.Errorf("Ticket not found")
	}

	student := repository.GetStudentByTicketId(ticket.Id)
	if student.Id == 0 {
		return fmt.Errorf("Student not found")
	}

	if student.ContractId != 0 {
		// so remove contract first

		contract := db.Contract{}
		err := repository.GetContractById(student.ContractId, &contract)
		if err != nil {
			return err
		}
		repository.RemoveContractById(contract.Id)
	}

	new_contract := repository.AddContract()

	repository.SetContract(student.Id, new_contract.Id)

	return nil
}

func Settle(student_ticket_number string, roomNumber string) error {
	student_ticket := repository.GetStudentTicketBySerialNumber(student_ticket_number)
	if student_ticket.Id == 0 {
		return fmt.Errorf("Ticket not found")
	}
	
	student := repository.GetStudentByTicketId(student_ticket.Id)
	if student.Id == 0 {
		return fmt.Errorf("Student not found")
	}

	if !repository.IsRoomNumberExists(roomNumber) {
		return fmt.Errorf("Room not found")
	}

	if student.ContractId == 0 {
		return fmt.Errorf("Student has no contract")
	}

	// verify contract
	contract := db.Contract{}
	err := repository.GetContractById(student.ContractId, &contract)
	if err != nil {
		return err
	}

	if contract.ExpireDate.Before(time.Now()) {
		return fmt.Errorf("Contract is expired")
	}

	room := repository.GetRoomByNumber(roomNumber)
	if room.Id == 0 {
		return fmt.Errorf("Room not found")
	}

	if room.IsMale != student.IsMale {
		return fmt.Errorf("Room is not suitable for this student")
	}

	places := repository.GetFreePlacesByRoomId(room.Id)
	if len(places) == 0 {
		return fmt.Errorf("No free places in room")
	}

	place := places[0]
	repository.SetStudentToPlace(student.Id, place.Id)

	return nil
}

func Unsettle(student_ticket_number string) error {
	student_ticket := repository.GetStudentTicketBySerialNumber(student_ticket_number)
	if student_ticket.Id == 0 {
		return fmt.Errorf("Ticket not found")
	}
	
	student := repository.GetStudentByTicketId(student_ticket.Id)
	if student.Id == 0 {
		return fmt.Errorf("Student not found")
	}

	if student.PlaceId == 0 {
		return fmt.Errorf("Student is not settled")
	}

	place_id := student.PlaceId
	place := db.Place{}
	err := repository.GetPlaceById(place_id, &place)
	if err != nil {
		return err
	}

	repository.UnsetStudentFromPlace(student.Id)

	return nil
}

func Resettle(student_ticket_number string, roomNumber string) error {
	err := Unsettle(student_ticket_number)
	if err != nil {
		return err
	}

	err = Settle(student_ticket_number, roomNumber)
	if err != nil {
		return err
	}

	return nil
}

func GetStudents(uuid string) (error, []db.Student) {
	if !IsUserAdmin(uuid) {
		return fmt.Errorf("User is not admin"), nil
	}

	return nil, repository.GetStudents()
}
