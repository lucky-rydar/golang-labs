package handlers

import (
	"encoding/json"
	"net/http"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode("Pong")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
