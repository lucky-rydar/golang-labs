package repository

import (
	"fmt"
	"time"

	"github.com/it-02/dormitory/db"
	"gorm.io/gorm"
)

type Contract struct {
	db *gorm.DB
}

func NewContract(db *gorm.DB) *Contract {
	return &Contract{db: db}
}

func (c *Contract) AddContract() db.Contract {
	contract := db.Contract{
		SignDate:   time.Now(),
		ExpireDate: time.Now().AddDate(1, 0, 0),
	}

	c.db.Create(&contract)
	fmt.Printf("Contract {id: %d} inserted\n", contract.Id)

	return contract
}

func (c *Contract) GetContracts() []db.Contract {
	var contracts []db.Contract
	c.db.Find(&contracts)
	return contracts
}

func (c *Contract) GetContractById(id uint, contract *db.Contract) error {
	var err error
	c.db.First(&contract, id)
	if contract.Id == 0 {
		err = fmt.Errorf("Contract with id %d not found", id)
	}
	return err
}

func (c *Contract) RemoveContractById(id uint) error {
	var err error
	var contract db.Contract
	err = c.GetContractById(id, &contract)
	if err == nil {
		c.db.Delete(&contract)
	}
	return err
}
