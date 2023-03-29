package service

import (
	"github.com/it-02/dormitory/repository"
	"github.com/it-02/dormitory/db"
)

func AddContract() db.Contract {
	return repository.AddContract()
}

func GetContracts() []db.Contract {
	return repository.GetContracts()
}
