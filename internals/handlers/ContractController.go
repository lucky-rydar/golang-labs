package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/it-02/dormitory/internals/db"
)

type IContractService interface {
	AddContract(uuid string) (db.Contract, error)
	GetContracts(uuid string) ([]db.Contract, error)
}

type ContractController struct {
	contract_service IContractService
}

func NewContractController(contract_service IContractService) *ContractController {
	return &ContractController{
		contract_service: contract_service,
	}
}

type AddContractRequest struct {
	UUID string `json:"uuid"`
}

func (cc *ContractController) AddContractHandler(w http.ResponseWriter, r *http.Request) {
	var request AddContractRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	contract, err := cc.contract_service.AddContract(request.UUID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(contract)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

type GetContractsRequest struct {
	UUID string `json:"uuid"`
}

func (cc *ContractController) GetContractsHandler(w http.ResponseWriter, r *http.Request) {
	var request GetContractsRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	contracts, err := cc.contract_service.GetContracts(request.UUID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(contracts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
