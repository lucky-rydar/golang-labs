package service

import (
	"fmt"

	"github.com/it-02/dormitory/repository"
	"github.com/it-02/dormitory/db"
)

type ContractService struct {
	contract_repository repository.IContract
	user_service IUserService
}

func NewContractService(contract_repository repository.IContract, user_service IUserService) *ContractService {
	return &ContractService{contract_repository: contract_repository, user_service: user_service}
}

func (cs *ContractService) AddContract(uuid string) (db.Contract, error) {
	if !cs.user_service.IsUserAdmin(uuid) {
		return db.Contract{}, fmt.Errorf("User is not admin")
	}

	return cs.contract_repository.AddContract(), nil
}

func (cs *ContractService) GetContracts(uuid string) ([]db.Contract, error) {
	if !cs.user_service.IsUserAdmin(uuid) {
		return []db.Contract{}, fmt.Errorf("User is not admin")
	}

	return cs.contract_repository.GetContracts(), nil
}
