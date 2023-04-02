package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/it-02/dormitory/service"
)

func GetPlacesHandler(w http.ResponseWriter, r *http.Request) {
	places := service.GetPlaces()
	err := json.NewEncoder(w).Encode(places)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func GetFreePlacesHandler(w http.ResponseWriter, r *http.Request) {
	places := service.GetFreePlaces()
	err := json.NewEncoder(w).Encode(places)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

type GetFreePlacesByRoomIdRequest struct {
	RoomId uint
}

func GetFreePlacesByRoomIdHandler(w http.ResponseWriter, r *http.Request) {
	var request GetFreePlacesByRoomIdRequest
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	places := service.GetFreePlacesByRoomId(request.RoomId)
	err = json.NewEncoder(w).Encode(places)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

type GetPlacesByRoomIdRequest struct {
	RoomId uint
}

func GetPlacesByRoomIdHandler(w http.ResponseWriter, r *http.Request) {
	var request GetPlacesByRoomIdRequest
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	places := service.GetPlacesByRoomId(request.RoomId)
	err = json.NewEncoder(w).Encode(places)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
