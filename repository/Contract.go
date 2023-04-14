package repository

import (
	"fmt"
	"time"

	"github.com/it-02/dormitory/db"
	"gorm.io/gorm"
)

type IConstract interface {
	AddContract() db.Contract
	GetContracts() []db.Contract
	GetContractById(id uint, contract *db.Contract) error
	RemoveContractById(id uint) error
}

type Contract struct {
	db *gorm.DB
}

func NewContract(db *gorm.DB) IConstract {
	return &Contract{db: db}
}

func (this Contract) AddContract() db.Contract {
	contract := db.Contract{
		SignDate:   time.Now(),
		ExpireDate: time.Now().AddDate(1, 0, 0),
	}

	this.db.Create(&contract)
	fmt.Printf("Contract {id: %d} inserted\n", contract.Id)

	return contract
}

func GetContracts() []db.Contract {
	var contracts []db.Contract
	this.db.Find(&contracts)
	return contracts
}

func GetContractById(id uint, contract *db.Contract) error {
	var err error
	this.db.First(&contract, id)
	if contract.Id == 0 {
		err = fmt.Errorf("Contract with id %d not found", id)
	}
	return err
}

func RemoveContractById(id uint) error {
	var err error
	var contract db.Contract
	err = GetContractById(id, &contract)
	if err == nil {
		this.db.Delete(&contract)
	}
	return err
}
