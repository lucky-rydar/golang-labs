package service

import (
	"fmt"
	"time"

	"github.com/it-02/dormitory/db"
	"github.com/it-02/dormitory/repository"
)

func AskAdminRegister(name string, surname string, isMale bool, studentTicketNumber string, studentTicketExpireDate time.Time) error {
	err := repository.AddRegisterAction(name, surname, isMale, studentTicketNumber, studentTicketExpireDate)
	if err != nil {
		return err
	}

	return nil
}

func AskAdminSignContract(studentTicketNumber string) error {
	err := repository.AddSignContractAction(studentTicketNumber)
	if err != nil {
		return err
	}

	return nil
}

func AskAdminUnsettle(studentTicketNumber string) error {
	err := repository.AddUnsettleAction(studentTicketNumber)
	if err != nil {
		return err
	}

	return nil
}

func AskAdminSettle(studentTicketNumber string, roomNumber string) error {
	err := repository.AddSettleAction(studentTicketNumber, roomNumber)
	if err != nil {
		return err
	}

	return nil
}

func AskAdminResettle(studentTicketNumber string, roomNumber string) error {
	err := repository.AddResettleAction(studentTicketNumber, roomNumber)
	if err != nil {
		return err
	}

	return nil
}

func GetActions(uuid string) ([]db.AskAdmin, error) {
	if !IsUserAdmin(uuid) {
		return nil, fmt.Errorf("User is not admin")
	}

	actions, err := repository.GetActions()
	if err != nil {
		return nil, err
	}

	return actions, nil
}

func ResolveAction(uuid string, actionId uint, isApproved bool) error {
	if !IsUserAdmin(uuid) {
		return fmt.Errorf("User is not admin")
	}

	if !isApproved {
		// just delete, no resolution is needed
		err := repository.DeleteActionById(actionId)
		if err != nil {
			return err
		}
		return nil
	}

	action, err := repository.GetActionById(actionId)
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

		err = RegisterStudent(&student, &student_ticket)
	} else if action.Action == "sign_contract" {
		err = SignContract(action.StudentTicketNumber)
	} else if action.Action == "unsettle" {
		err = Unsettle(action.StudentTicketNumber)
	} else if action.Action == "settle" {
		err = Settle(action.StudentTicketNumber, action.RoomNumber)
	} else if action.Action == "resettle" {
		err = Resettle(action.StudentTicketNumber, action.RoomNumber)
	} else {
		err = fmt.Errorf("Unknown action")
	}

	// remove action 
	remove_err := repository.DeleteActionById(actionId)
	if remove_err != nil {
		return remove_err
	}

	if err != nil {
		return err
	}

	return nil
}
