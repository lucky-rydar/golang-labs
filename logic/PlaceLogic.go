package logic

import (
	"fmt"

	"github.com/it-02/dormitory/db"
	"github.com/it-02/dormitory/models"
)

func GetPlaces() []models.Place {
	var places []models.Place
	db.DB.Find(&places)
	return places
}

func GetFreePlaces() []models.Place {
	var places []models.Place
	db.DB.Where("is_free = ?", true).Find(&places)
	return places
}

func GetFreePlacesByRoomId(roomId uint) []models.Place {
	var places []models.Place
	db.DB.Where("is_free = ? AND room_id = ?", true, roomId).Find(&places)
	return places
}

func GetPlacesByRoomId(roomId uint) []models.Place {
	var places []models.Place
	db.DB.Where("room_id = ?", roomId).Find(&places)
	return places
}

func GetPlaceById(id uint, place *models.Place) error {
	var err error
	db.DB.First(&place, id)
	if place.Id == 0 {
		err = fmt.Errorf("Place with id %d not found", id)
	}
	return err
}

func GetPlacesByParams(isMale bool, isFree bool) []models.Place {
	var places []models.Place
	db.DB.Where("is_free = ?", isFree).Find(&places)

	var placesRet []models.Place
	for i := 0; i < len(places); i++ {
		var room models.Room
		db.DB.First(&room, places[i].RoomId)
		if room.IsMale == isMale {
			placesRet = append(placesRet, places[i])
		}
	}

	return placesRet
}
