package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/it-02/dormitory/service"
)

type PlaceController struct {
	place_service service.IPlaceService
}

func NewPlaceController(place_service service.IPlaceService) *PlaceController {
	return &PlaceController{
		place_service: place_service,
	}
}

func (pc *PlaceController) GetPlacesHandler(w http.ResponseWriter, r *http.Request) {
	places := pc.place_service.GetPlaces()
	err := json.NewEncoder(w).Encode(places)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (pc *PlaceController) GetFreePlacesHandler(w http.ResponseWriter, r *http.Request) {
	places := pc.place_service.GetFreePlaces()
	err := json.NewEncoder(w).Encode(places)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

type GetFreePlacesByRoomIdRequest struct {
	RoomId uint
}

func (pc *PlaceController) GetFreePlacesByRoomIdHandler(w http.ResponseWriter, r *http.Request) {
	var request GetFreePlacesByRoomIdRequest
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	places := pc.place_service.GetFreePlacesByRoomId(request.RoomId)
	err = json.NewEncoder(w).Encode(places)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

type GetPlacesByRoomIdRequest struct {
	RoomId uint
}

func (pc *PlaceController) GetPlacesByRoomIdHandler(w http.ResponseWriter, r *http.Request) {
	var request GetPlacesByRoomIdRequest
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	places := pc.place_service.GetPlacesByRoomId(request.RoomId)
	err = json.NewEncoder(w).Encode(places)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
