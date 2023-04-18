package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/it-02/dormitory/service"
	"github.com/it-02/dormitory/db"
)

type IRoomService interface {
	AddRoom(uuid string, room *db.Room) error
	GetRooms() []db.Room
	GetRoomByPlaceId(placeId uint) db.Room
	GetRoomStatsByNumber(number string, room_stats *service.RoomStats) error
}

type RoomController struct {
	room_service IRoomService
}

func NewRoomController(room_service IRoomService) *RoomController {
	return &RoomController{
		room_service: room_service,
	}
}

type AddRoomRequest struct {
	IsMale       bool
	AreaSqMeters float32
	Number 	     string
	UUID 	     string `json:"uuid"`
}

func (rc *RoomController) AddRoomHandler(w http.ResponseWriter, r *http.Request) {
	var request AddRoomRequest
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	room := db.Room{
		IsMale:       request.IsMale,
		AreaSqMeters: request.AreaSqMeters,
		Number:       request.Number,
	}

	err = rc.room_service.AddRoom(request.UUID, &room)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

// has no request body
func (rc *RoomController) GetRoomsHandler(w http.ResponseWriter, r *http.Request) {
	rooms := rc.room_service.GetRooms()
	err := json.NewEncoder(w).Encode(rooms)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

type GetRoomByPlaceId struct {
	PlaceId uint
}

func (rc *RoomController) GetRoomByPlaceIdHandler(w http.ResponseWriter, r *http.Request) {
	var request GetRoomByPlaceId
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	room := rc.room_service.GetRoomByPlaceId(request.PlaceId)
	err = json.NewEncoder(w).Encode(room)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

type GetRoomStatsByNumberRequest struct {
	Number string
}

func (rc *RoomController) GetRoomStatsByNumberHandler(w http.ResponseWriter, r *http.Request) {
	var request GetRoomStatsByNumberRequest
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var roomStats service.RoomStats
	err = rc.room_service.GetRoomStatsByNumber(request.Number, &roomStats)
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
