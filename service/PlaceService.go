package service

import (
	"github.com/it-02/dormitory/repository"
	"github.com/it-02/dormitory/db"
)

func GetPlaces() []db.Place {
	return repository.GetPlaces()
}

func GetFreePlaces() []db.Place {
	return repository.GetFreePlaces()
}

func GetFreePlacesByRoomId(roomId uint) []db.Place {
	return repository.GetFreePlacesByRoomId(roomId)
}

func GetPlacesByRoomId(roomId uint) []db.Place {
	return repository.GetPlacesByRoomId(roomId)
}
