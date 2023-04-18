package service

import (
	"fmt"
	"time"

	"github.com/it-02/dormitory/db"
	"github.com/it-02/dormitory/repository"
)

type IStudentService interface {
	RegisterStudent(student *db.Student, student_ticket *db.StudentTicket) error
	SignContract(student_ticket_number string) error
	Settle(student_ticket_number string, roomNumber string) error
	Unsettle(student_ticket_number string) error
	Resettle(student_ticket_number string, roomNumber string) error
	GetStudents(uuid string) (error, []StudentRepr)
}

type StudentService struct {
	student_repository repository.IStudent
	student_ticket_repository repository.IStudentTicket
	room_repository repository.IRoom
	place_repository repository.IPlace
	contract_repository repository.IContract
	user_service IUserService
}

func NewStudentService(student_repository repository.IStudent, student_ticket_repository repository.IStudentTicket, room_repository repository.IRoom, place_repository repository.IPlace, contract_repository repository.IContract, user_service IUserService) *StudentService {
	return &StudentService{student_repository: student_repository, student_ticket_repository: student_ticket_repository, room_repository: room_repository, place_repository: place_repository, contract_repository: contract_repository, user_service: user_service}
}

func (ss *StudentService) RegisterStudent(student *db.Student, student_ticket *db.StudentTicket) error {
	ret := ss.student_ticket_repository.AddStudentTicket(student_ticket)
	if ret != nil {
		return ret
	}

	student.StudentTicketId = student_ticket.Id
	ret = ss.student_repository.AddStudent(student)
	if ret != nil {
		return ret
	}

	return ret
}

func (ss *StudentService) SignContract(student_ticket_number string) error {
	ticket := ss.student_ticket_repository.GetStudentTicketBySerialNumber(student_ticket_number)
	if ticket.Id == 0 {
		return fmt.Errorf("Ticket not found")
	}

	student := ss.student_repository.GetStudentByTicketId(ticket.Id)
	if student.Id == 0 {
		return fmt.Errorf("Student not found")
	}

	if student.ContractId != 0 {
		// so remove contract first

		contract := db.Contract{}
		err := ss.contract_repository.GetContractById(student.ContractId, &contract)
		if err != nil {
			return err
		}
		err = ss.contract_repository.RemoveContractById(contract.Id)
		if err != nil {
			return err
		}
	}

	new_contract := ss.contract_repository.AddContract()

	err := ss.student_repository.SetContract(student.Id, new_contract.Id)
	if err != nil {
		return err
	}

	return nil
}

func (ss *StudentService) Settle(student_ticket_number string, roomNumber string) error {
	student_ticket := ss.student_ticket_repository.GetStudentTicketBySerialNumber(student_ticket_number)
	if student_ticket.Id == 0 {
		return fmt.Errorf("Ticket not found")
	}
	
	student := ss.student_repository.GetStudentByTicketId(student_ticket.Id)
	if student.Id == 0 {
		return fmt.Errorf("Student not found")
	}

	if !ss.room_repository.IsRoomNumberExists(roomNumber) {
		return fmt.Errorf("Room not found")
	}

	if student.ContractId == 0 {
		return fmt.Errorf("Student has no contract")
	}

	// verify contract
	contract := db.Contract{}
	err := ss.contract_repository.GetContractById(student.ContractId, &contract)
	if err != nil {
		return err
	}

	if contract.ExpireDate.Before(time.Now()) {
		return fmt.Errorf("Contract is expired")
	}

	room := ss.room_repository.GetRoomByNumber(roomNumber)
	if room.Id == 0 {
		return fmt.Errorf("Room not found")
	}

	if room.IsMale != student.IsMale {
		return fmt.Errorf("Room is not suitable for ss student")
	}

	places := ss.place_repository.GetFreePlacesByRoomId(room.Id)
	if len(places) == 0 {
		return fmt.Errorf("No free places in room")
	}

	place := places[0]
	err = ss.student_repository.SetStudentToPlace(student.Id, place.Id)
	if err != nil {
		return err
	}

	return nil
}

func (ss *StudentService) Unsettle(student_ticket_number string) error {
	student_ticket := ss.student_ticket_repository.GetStudentTicketBySerialNumber(student_ticket_number)
	if student_ticket.Id == 0 {
		return fmt.Errorf("Ticket not found")
	}
	
	student := ss.student_repository.GetStudentByTicketId(student_ticket.Id)
	if student.Id == 0 {
		return fmt.Errorf("Student not found")
	}

	if student.PlaceId == 0 {
		return fmt.Errorf("Student is not settled")
	}

	place_id := student.PlaceId
	place := db.Place{}
	err := ss.place_repository.GetPlaceById(place_id, &place)
	if err != nil {
		return err
	}

	err = ss.student_repository.UnsetStudentFromPlace(student.Id)
	if err != nil {
		return err
	}

	return nil
}

func (ss *StudentService) Resettle(student_ticket_number string, roomNumber string) error {
	err := ss.Unsettle(student_ticket_number)
	if err != nil {
		return err
	}

	err = ss.Settle(student_ticket_number, roomNumber)
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

func (ss *StudentService) GetStudents(uuid string) (error, []StudentRepr) {
	if !ss.user_service.IsUserAdmin(uuid) {
		return fmt.Errorf("User is not admin"), nil
	}

	ret := []StudentRepr{};

	students := ss.student_repository.GetStudents()
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
			err := ss.contract_repository.GetContractById(student.ContractId, &contract)
			if err != nil {
				return err, nil
			}
			student_repr.Contract = contract
		}

		if student.StudentTicketId != 0 {
			student_ticket := ss.student_ticket_repository.GetStudentTicketById(student.StudentTicketId)
			if student_ticket.Id == 0 {
				// error should be returned because student can't be registered without a ticket
				return fmt.Errorf("Ticket not found"), nil
			}
			student_repr.StudentTicket = student_ticket
		}

		if student.PlaceId != 0 {
			place := db.Place{}
			err := ss.place_repository.GetPlaceById(student.PlaceId, &place)
			if err != nil {
				return err, nil
			}
			
			room := db.Room{}
			err = ss.room_repository.GetRoomById(place.RoomId, &room)
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
