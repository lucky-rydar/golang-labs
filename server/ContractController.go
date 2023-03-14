package server

import (
	"time"
	"encoding/json"
	"net/http"

	"github.com/it-02/dormitory/logic"
	"github.com/it-02/dormitory/models"
)

func AddContractHandler(w http.ResponseWriter, r *http.Request) {
	var contract models.Contract

	contract.SignDate = time.Now()

	// expire date is 1 year after sign date
	contract.ExpireDate = contract.SignDate.AddDate(1, 0, 0)

	logic.AddContract(&contract)
	err := json.NewEncoder(w).Encode(contract)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func GetContractsHandler(w http.ResponseWriter, r *http.Request) {
	contracts := logic.GetContracts()
	err := json.NewEncoder(w).Encode(contracts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
