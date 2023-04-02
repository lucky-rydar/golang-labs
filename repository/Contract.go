package repository

import (
	"fmt"
	"time"

	"github.com/it-02/dormitory/db"
)

func AddContract() db.Contract {
	contract := db.Contract{
		SignDate:   time.Now(),
		ExpireDate: time.Now().AddDate(1, 0, 0),
	}

	db.DB.Create(&contract)
	fmt.Printf("Contract {id: %d} inserted\n", contract.Id)

	return contract
}

func GetContracts() []db.Contract {
	var contracts []db.Contract
	db.DB.Find(&contracts)
	return contracts
}

func GetContractById(id uint, contract *db.Contract) error {
	var err error
	db.DB.First(&contract, id)
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
		db.DB.Delete(&contract)
	}
	return err
}
