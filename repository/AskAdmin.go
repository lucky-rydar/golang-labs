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

func (this AskAdmin) AddRegisterAction(name string, surname string, isMale bool, studentTicketNumber string, studentTicketExpireDate time.Time) error {
	var action db.AskAdmin
	action.Action = "register"
	action.Name = name
	action.Surname = surname
	action.IsMale = isMale
	action.StudentTicketNumber = studentTicketNumber
	action.StudentTicketExpireDate = studentTicketExpireDate

	result := this.db.Create(&action)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (this AskAdmin) AddSignContractAction(studentTicketNumber string) error {
	var action db.AskAdmin
	action.Action = "sign_contract"
	action.StudentTicketNumber = studentTicketNumber

	result := this.db.Create(&action)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (this AskAdmin) AddUnsettleAction(studentTicketNumber string) error {
	var action db.AskAdmin
	action.Action = "unsettle"
	action.StudentTicketNumber = studentTicketNumber

	result := this.db.Create(&action)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (this AskAdmin) AddSettleAction(studentTicketNumber string, roomNumber string) error {
	var action db.AskAdmin
	action.Action = "settle"
	action.StudentTicketNumber = studentTicketNumber
	action.RoomNumber = roomNumber

	result := this.db.Create(&action)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (this AskAdmin) AddResettleAction(studentTicketNumber string, roomNumber string) error {
	var action db.AskAdmin
	action.Action = "resettle"
	action.StudentTicketNumber = studentTicketNumber
	action.RoomNumber = roomNumber

	result := this.db.Create(&action)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (this AskAdmin) GetActions() ([]db.AskAdmin, error) {
	var actions []db.AskAdmin
	result := this.db.Find(&actions)
	if result.Error != nil {
		return nil, result.Error
	}

	return actions, nil
}

func (this AskAdmin) GetActionById(id uint) (db.AskAdmin, error) {
	var action db.AskAdmin
	result := this.db.First(&action, id)
	if result.Error != nil {
		return db.AskAdmin{}, result.Error
	}

	return action, nil
}

func (this AskAdmin) DeleteActionById(id uint) error {
	result := this.db.Delete(&db.AskAdmin{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
