package service

import (
	"github.com/it-02/dormitory/repository"
	"github.com/it-02/dormitory/db"
)

type IPlaceService interface {
	GetPlaces() []db.Place
	GetFreePlaces() []db.Place
	GetFreePlacesByRoomId(roomId uint) []db.Place
	GetPlacesByRoomId(roomId uint) []db.Place
}

type PlaceService struct {
	place_repository *repository.IPlace
}

func NewPlaceService(place_repository *repository.IPlace) IPlaceService {
	return &PlaceService{place_repository: place_repository}
}

func (this PlaceService) GetPlaces() []db.Place {
	return repository.GetPlaces()
}

func (this PlaceService) GetFreePlaces() []db.Place {
	return repository.GetFreePlaces()
}

func (this PlaceService) GetFreePlacesByRoomId(roomId uint) []db.Place {
	return repository.GetFreePlacesByRoomId(roomId)
}

func (this PlaceService) GetPlacesByRoomId(roomId uint) []db.Place {
	return repository.GetPlacesByRoomId(roomId)
}
