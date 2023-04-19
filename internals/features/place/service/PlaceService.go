package service

import (
	"github.com/it-02/dormitory/repository"
	"github.com/it-02/dormitory/internals/db"
)

type IPlace interface {
	GetPlaces() []db.Place
	GetFreePlaces() []db.Place
	GetFreePlacesByRoomId(roomId uint) []db.Place
	GetOccupiedPlacesByRoomId(roomId uint) []db.Place
	GetPlacesByRoomId(roomId uint) []db.Place
	GetPlaceById(id uint, place *db.Place) error
	GetPlacesByParams(isMale bool, isFree bool) []db.Place
}

type PlaceService struct {
	place_repository IPlace
}

func NewPlaceService(place_repository repository.IPlace) *PlaceService {
	return &PlaceService{place_repository: place_repository}
}

func (ps *PlaceService) GetPlaces() []db.Place {
	return ps.place_repository.GetPlaces()
}

func (ps *PlaceService) GetFreePlaces() []db.Place {
	return ps.place_repository.GetFreePlaces()
}

func (ps *PlaceService) GetFreePlacesByRoomId(roomId uint) []db.Place {
	return ps.place_repository.GetFreePlacesByRoomId(roomId)
}

func (ps *PlaceService) GetPlacesByRoomId(roomId uint) []db.Place {
	return ps.place_repository.GetPlacesByRoomId(roomId)
}
