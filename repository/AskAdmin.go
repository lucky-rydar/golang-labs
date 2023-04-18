package repository

import (
	"time"

	"github.com/it-02/dormitory/db"
	"gorm.io/gorm"
)

type IAskAdmin interface {
	AddRegisterAction(name string, surname string, isMale bool, studentTicketNumber string, studentTicketExpireDate time.Time) error
	AddSignContractAction(studentTicketNumber string) error
	AddUnsettleAction(studentTicketNumber string) error
	AddSettleAction(studentTicketNumber string, roomNumber string) error
	AddResettleAction(studentTicketNumber string, roomNumber string) error
	GetActions() ([]db.AskAdmin, error)
	GetActionById(id uint) (db.AskAdmin, error)
	DeleteActionById(id uint) error
}

type AskAdmin struct {
	db *gorm.DB
}

func NewAskAdmin(db *gorm.DB) IAskAdmin {
	return &AskAdmin{db: db}
}

func (aa *AskAdmin) AddRegisterAction(name string, surname string, isMale bool, studentTicketNumber string, studentTicketExpireDate time.Time) error {
	var action db.AskAdmin
	action.Action = "register"
	action.Name = name
	action.Surname = surname
	action.IsMale = isMale
	action.StudentTicketNumber = studentTicketNumber
	action.StudentTicketExpireDate = studentTicketExpireDate

	result := aa.db.Create(&action)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (aa *AskAdmin) AddSignContractAction(studentTicketNumber string) error {
	var action db.AskAdmin
	action.Action = "sign_contract"
	action.StudentTicketNumber = studentTicketNumber

	result := aa.db.Create(&action)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (aa *AskAdmin) AddUnsettleAction(studentTicketNumber string) error {
	var action db.AskAdmin
	action.Action = "unsettle"
	action.StudentTicketNumber = studentTicketNumber

	result := aa.db.Create(&action)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (aa *AskAdmin) AddSettleAction(studentTicketNumber string, roomNumber string) error {
	var action db.AskAdmin
	action.Action = "settle"
	action.StudentTicketNumber = studentTicketNumber
	action.RoomNumber = roomNumber

	result := aa.db.Create(&action)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (aa *AskAdmin) AddResettleAction(studentTicketNumber string, roomNumber string) error {
	var action db.AskAdmin
	action.Action = "resettle"
	action.StudentTicketNumber = studentTicketNumber
	action.RoomNumber = roomNumber

	result := aa.db.Create(&action)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (aa *AskAdmin) GetActions() ([]db.AskAdmin, error) {
	var actions []db.AskAdmin
	result := aa.db.Find(&actions)
	if result.Error != nil {
		return nil, result.Error
	}

	return actions, nil
}

func (aa *AskAdmin) GetActionById(id uint) (db.AskAdmin, error) {
	var action db.AskAdmin
	result := aa.db.First(&action, id)
	if result.Error != nil {
		return db.AskAdmin{}, result.Error
	}

	return action, nil
}

func (aa *AskAdmin) DeleteActionById(id uint) error {
	result := aa.db.Delete(&db.AskAdmin{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
