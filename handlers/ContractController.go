package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/it-02/dormitory/service"
)

type IContractController interface {
	AddContractHandler(w http.ResponseWriter, r *http.Request)
	GetContractsHandler(w http.ResponseWriter, r *http.Request)
}

type ContractController struct {
	contract_service *this.contract_service.IContractService
}

func NewContractController(contract_service *this.contract_service.IContractService) *ContractController {
	return &ContractController{
		contract_service: contract_service,
	}
}

type AddContractRequest struct {
	UUID string `json:"uuid"`
}

func (this ContractController) AddContractHandler(w http.ResponseWriter, r *http.Request) {
	var request AddContractRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	contract, err := this.contract_service.AddContract(request.UUID)
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

func (this ContractController) GetContractsHandler(w http.ResponseWriter, r *http.Request) {
	var request GetContractsRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	contracts, err := this.contract_service.GetContracts(request.UUID)
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
