package service

import (
	"fmt"

	"github.com/it-02/dormitory/repository"
	"github.com/it-02/dormitory/db"
)

func AddContract(uuid string) (db.Contract, error) {
	if !IsUserAdmin(uuid) {
		return db.Contract{}, fmt.Errorf("User is not admin")
	}

	return repository.AddContract(), nil
}

func GetContracts(uuid string) ([]db.Contract, error) {
	if !IsUserAdmin(uuid) {
		return []db.Contract{}, fmt.Errorf("User is not admin")
	}

	return repository.GetContracts(), nil
}
