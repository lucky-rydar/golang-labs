package repository

import (
	"time"

	"github.com/it-02/dormitory/db"
)

func AddRegisterAction(name string, surname string, isMale bool, studentTicketNumber string, studentTicketExpireDate time.Time) error {
	var action db.AskAdmin
	action.Action = "register"
	action.Name = name
	action.Surname = surname
	action.IsMale = isMale
	action.StudentTicketNumber = studentTicketNumber
	action.StudentTicketExpireDate = studentTicketExpireDate

	result := db.DB.Create(&action)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func AddSignContractAction(studentTicketNumber string) error {
	var action db.AskAdmin
	action.Action = "sign_contract"
	action.StudentTicketNumber = studentTicketNumber

	result := db.DB.Create(&action)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func AddUnsettleAction(studentTicketNumber string) error {
	var action db.AskAdmin
	action.Action = "unsettle"
	action.StudentTicketNumber = studentTicketNumber

	result := db.DB.Create(&action)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func AddSettleAction(studentTicketNumber string, roomNumber string) error {
	var action db.AskAdmin
	action.Action = "settle"
	action.StudentTicketNumber = studentTicketNumber
	action.RoomNumber = roomNumber

	result := db.DB.Create(&action)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func AddResettleAction(studentTicketNumber string, roomNumber string) error {
	var action db.AskAdmin
	action.Action = "resettle"
	action.StudentTicketNumber = studentTicketNumber
	action.RoomNumber = roomNumber

	result := db.DB.Create(&action)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func GetActions() ([]db.AskAdmin, error) {
	var actions []db.AskAdmin
	result := db.DB.Find(&actions)
	if result.Error != nil {
		return nil, result.Error
	}

	return actions, nil
}

func GetActionById(id uint) (db.AskAdmin, error) {
	var action db.AskAdmin
	result := db.DB.First(&action, id)
	if result.Error != nil {
		return db.AskAdmin{}, result.Error
	}

	return action, nil
}

func DeleteActionById(id uint) error {
	result := db.DB.Delete(&db.AskAdmin{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
