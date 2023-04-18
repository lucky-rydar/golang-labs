package service

import (
	"fmt"
	"time"

	"github.com/it-02/dormitory/db"
	"github.com/it-02/dormitory/repository"
)

type IAskAdminService interface {
	AskAdminRegister(name string, surname string, isMale bool, studentTicketNumber string, studentTicketExpireDate time.Time) error
	AskAdminSignContract(studentTicketNumber string) error
	AskAdminUnsettle(studentTicketNumber string) error
	AskAdminSettle(studentTicketNumber string, roomNumber string) error
	AskAdminResettle(studentTicketNumber string, roomNumber string) error
	GetActions(uuid string) ([]db.AskAdmin, error)
	ResolveAction(uuid string, actionId uint, isApproved bool) error
}

type AskAdminService struct {
	ask_admin_repository repository.IAskAdmin
	user_service IUserService
	student_service IStudentService
}

func NewAskAdminService(ask_admin_repository repository.IAskAdmin, user_service IUserService, student_service IStudentService) IAskAdminService {
	return &AskAdminService{ask_admin_repository: ask_admin_repository, user_service: user_service, student_service: student_service}
}

func (aas *AskAdminService) AskAdminRegister(name string, surname string, isMale bool, studentTicketNumber string, studentTicketExpireDate time.Time) error {
	err := aas.ask_admin_repository.AddRegisterAction(name, surname, isMale, studentTicketNumber, studentTicketExpireDate)
	if err != nil {
		return err
	}

	return nil
}

func (aas *AskAdminService) AskAdminSignContract(studentTicketNumber string) error {
	err := aas.ask_admin_repository.AddSignContractAction(studentTicketNumber)
	if err != nil {
		return err
	}

	return nil
}

func (aas *AskAdminService) AskAdminUnsettle(studentTicketNumber string) error {
	err := aas.ask_admin_repository.AddUnsettleAction(studentTicketNumber)
	if err != nil {
		return err
	}

	return nil
}

func (aas *AskAdminService) AskAdminSettle(studentTicketNumber string, roomNumber string) error {
	err := aas.ask_admin_repository.AddSettleAction(studentTicketNumber, roomNumber)
	if err != nil {
		return err
	}

	return nil
}

func (aas *AskAdminService) AskAdminResettle(studentTicketNumber string, roomNumber string) error {
	err := aas.ask_admin_repository.AddResettleAction(studentTicketNumber, roomNumber)
	if err != nil {
		return err
	}

	return nil
}

func (aas *AskAdminService) GetActions(uuid string) ([]db.AskAdmin, error) {
	if !aas.user_service.IsUserAdmin(uuid) {
		return nil, fmt.Errorf("User is not admin")
	}

	actions, err := aas.ask_admin_repository.GetActions()
	if err != nil {
		return nil, err
	}

	return actions, nil
}

func (aas *AskAdminService) ResolveAction(uuid string, actionId uint, isApproved bool) error {
	if !aas.user_service.IsUserAdmin(uuid) {
		return fmt.Errorf("User is not admin")
	}

	if !isApproved {
		// just delete, no resolution is needed
		err := aas.ask_admin_repository.DeleteActionById(actionId)
		if err != nil {
			return err
		}
		return nil
	}

	action, err := aas.ask_admin_repository.GetActionById(actionId)
	if err != nil {
		return err
	}

	if action.Action == "register" {
		student := db.Student{
			Name:                    action.Name,
			Surname:                 action.Surname,
			IsMale:                  action.IsMale,
			ContractId:              0,
			StudentTicketId:         0,
			PlaceId:                 0,
		}

		student_ticket := db.StudentTicket{
			SerialNumber:     action.StudentTicketNumber,
			ExpireDate: action.StudentTicketExpireDate,
		}

		err = aas.student_service.RegisterStudent(&student, &student_ticket)
	} else if action.Action == "sign_contract" {
		err = aas.student_service.SignContract(action.StudentTicketNumber)
	} else if action.Action == "unsettle" {
		err = aas.student_service.Unsettle(action.StudentTicketNumber)
	} else if action.Action == "settle" {
		err = aas.student_service.Settle(action.StudentTicketNumber, action.RoomNumber)
	} else if action.Action == "resettle" {
		err = aas.student_service.Resettle(action.StudentTicketNumber, action.RoomNumber)
	} else {
		err = fmt.Errorf("Unknown action")
	}

	// remove action 
	remove_err := aas.ask_admin_repository.DeleteActionById(actionId)
	if remove_err != nil {
		return remove_err
	}

	if err != nil {
		return err
	}

	return nil
}
