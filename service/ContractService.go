package service

import (
	"fmt"

	"github.com/it-02/dormitory/repository"
	"github.com/it-02/dormitory/db"
)

type IContractService interface {
	AddContract(uuid string) (db.Contract, error)
	GetContracts(uuid string) ([]db.Contract, error)
}

type ContractService struct {
	contract_repository repository.IContract
	user_service IUserService
}

func NewContractService(contract_repository repository.IContract, user_service IUserService) IContractService {
	return &ContractService{contract_repository: contract_repository, user_service: user_service}
}

func (this ContractService) AddContract(uuid string) (db.Contract, error) {
	if !this.user_service.IsUserAdmin(uuid) {
		return db.Contract{}, fmt.Errorf("User is not admin")
	}

	return this.contract_repository.AddContract(), nil
}

func (this ContractService) GetContracts(uuid string) ([]db.Contract, error) {
	if !this.user_service.IsUserAdmin(uuid) {
		return []db.Contract{}, fmt.Errorf("User is not admin")
	}

	return this.contract_repository.GetContracts(), nil
}
