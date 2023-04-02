package repository

import (
	"fmt"

	"github.com/it-02/dormitory/db"
)

func GetPlaces() []db.Place {
	var places []db.Place
	db.DB.Find(&places)
	return places
}

func GetFreePlaces() []db.Place {
	var places []db.Place
	db.DB.Where("is_free = ?", true).Find(&places)
	return places
}

func GetFreePlacesByRoomId(roomId uint) []db.Place {
	var places []db.Place
	db.DB.Where("is_free = ? AND room_id = ?", true, roomId).Find(&places)
	return places
}

func GetPlacesByRoomId(roomId uint) []db.Place {
	var places []db.Place
	db.DB.Where("room_id = ?", roomId).Find(&places)
	return places
}

func GetPlaceById(id uint, place *db.Place) error {
	var err error
	db.DB.First(&place, id)
	if place.Id == 0 {
		err = fmt.Errorf("Place with id %d not found", id)
	}
	return err
}

func GetPlacesByParams(isMale bool, isFree bool) []db.Place {
	var places []db.Place
	db.DB.Where("is_free = ?", isFree).Find(&places)

	var placesRet []db.Place
	for i := 0; i < len(places); i++ {
		var room db.Room
		db.DB.First(&room, places[i].RoomId)
		if room.IsMale == isMale {
			placesRet = append(placesRet, places[i])
		}
	}

	return placesRet
}
