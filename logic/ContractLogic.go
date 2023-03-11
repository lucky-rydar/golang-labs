package logic

import (
	"fmt"

	"github.com/it-02/dormitory/db"
	"github.com/it-02/dormitory/models"
)

func AddContract(contract *models.Contract) {
	db.DB.Create(&contract)
	fmt.Printf("Contract {id: %d} inserted\n", contract.Id)
}

func GetContracts() []models.Contract {
	var contracts []models.Contract
	db.DB.Find(&contracts)
	return contracts
}
