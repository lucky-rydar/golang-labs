package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/it-02/dormitory/service"
)

func AddContractHandler(w http.ResponseWriter, r *http.Request) {
	contract := service.AddContract()
	err := json.NewEncoder(w).Encode(contract)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func GetContractsHandler(w http.ResponseWriter, r *http.Request) {
	contracts := service.GetContracts()
	err := json.NewEncoder(w).Encode(contracts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
