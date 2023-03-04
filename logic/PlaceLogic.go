package logic

import (
	"fmt"

	"github.com/it-02/dormitory/db"
	"github.com/it-02/dormitory/models"
)

func AddPlace(place models.Place) error {
	var ret error
	ret = nil

	var room models.Room
	db.DB.First(&room, place.RoomId)
	if room.Id != 0 {
		db.DB.Create(&place)
	} else {
		ret = fmt.Errorf("Room not found")
		fmt.Printf("Room not found")
	}

	return ret
}

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

