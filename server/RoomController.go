package server

import (
	"encoding/json"
	"net/http"

	"github.com/it-02/dormitory/logic"
	"github.com/it-02/dormitory/models"
)

func AddRoomHandler(w http.ResponseWriter, r *http.Request) {
	var room models.Room
	var err error

	err = json.NewDecoder(r.Body).Decode(&room)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	logic.AddRoom(room)

	err = json.NewEncoder(w).Encode(room)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func GetRoomsHandler(w http.ResponseWriter, r *http.Request) {
	rooms := logic.GetRooms()
	err := json.NewEncoder(w).Encode(rooms)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
