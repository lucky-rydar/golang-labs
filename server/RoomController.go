package server

import (
	"encoding/json"
	"net/http"

	"github.com/it-02/dormitory/logic"
	"github.com/it-02/dormitory/models"
)

type AddRoomRequest struct {
	IsMale       bool
	AreaSqMeters float32
	Number 	     string
}

func AddRoomHandler(w http.ResponseWriter, r *http.Request) {
	var request AddRoomRequest
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	room := models.Room{
		IsMale:       request.IsMale,
		AreaSqMeters: request.AreaSqMeters,
		Number:       request.Number,
	}

	logic.AddRoom(room)
}

// has no request body
func GetRoomsHandler(w http.ResponseWriter, r *http.Request) {
	rooms := logic.GetRooms()
	err := json.NewEncoder(w).Encode(rooms)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

type GetRoomByPlaceId struct {
	PlaceId uint
}

func GetRoomByPlaceIdHandler(w http.ResponseWriter, r *http.Request) {
	var request GetRoomByPlaceId
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	room := logic.GetRoomByPlaceId(request.PlaceId)
	err = json.NewEncoder(w).Encode(room)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

type GetRoomStatsByNumberRequest struct {
	Number string
}

func GetRoomStatsByNumberHandler(w http.ResponseWriter, r *http.Request) {
	var request GetRoomStatsByNumberRequest
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var roomStats logic.RoomStats
	err = logic.GetRoomStatsByNumber(request.Number, &roomStats)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = json.NewEncoder(w).Encode(roomStats)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
