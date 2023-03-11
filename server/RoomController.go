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
