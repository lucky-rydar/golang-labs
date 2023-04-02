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
		err = repository.RemoveContractById(contract.Id)
		if err != nil {
			return err
		}
	}

	new_contract := repository.AddContract()

	err := repository.SetContract(student.Id, new_contract.Id)
	if err != nil {
		return err
	}

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
	err = repository.SetStudentToPlace(student.Id, place.Id)
	if err != nil {
		return err
	}

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

	err = repository.UnsetStudentFromPlace(student.Id)
	if err != nil {
		return err
	}

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

type StudentRepr struct {
	Id          uint
	Name        string
	Surname     string
	IsMale	    bool
	Place       PlaceRepr
	Contract    db.Contract
	StudentTicket db.StudentTicket
}

func GetStudents(uuid string) (error, []StudentRepr) {
	if !IsUserAdmin(uuid) {
		return fmt.Errorf("User is not admin"), nil
	}

	ret := []StudentRepr{};

	students := repository.GetStudents()
	for i := 0; i < len(students); i++ {
		student := students[i]

		student_repr := StudentRepr{
			Id: student.Id,
			Name: student.Name,
			Surname: student.Surname,
			IsMale: student.IsMale,
		}

		if student.ContractId != 0 {
			contract := db.Contract{}
			err := repository.GetContractById(student.ContractId, &contract)
			if err != nil {
				return err, nil
			}
			student_repr.Contract = contract
		}

		if student.StudentTicketId != 0 {
			student_ticket := repository.GetStudentTicketById(student.StudentTicketId)
			if student_ticket.Id == 0 {
				// error should be returned because student can't be registered without a ticket
				return fmt.Errorf("Ticket not found"), nil
			}
			student_repr.StudentTicket = student_ticket
		}

		if student.PlaceId != 0 {
			place := db.Place{}
			err := repository.GetPlaceById(student.PlaceId, &place)
			if err != nil {
				return err, nil
			}
			
			room := db.Room{}
			err = repository.GetRoomById(place.RoomId, &room)
			if err != nil {
				return err, nil
			}

			place_repr := PlaceRepr{
				PlaceId: place.Id,
				IsFree: place.IsFree,
				IsMale: room.IsMale,
				RoomNumber: room.Number,
			}

			student_repr.Place = place_repr
		}

		ret = append(ret, student_repr)
	}

	return nil, ret
}
