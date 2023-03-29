package service

import (
	"fmt"

	"github.com/it-02/dormitory/repository"
	"github.com/it-02/dormitory/db"
)

func AddRoom(room db.Room) {
	repository.AddRoom(room)
}

func GetRooms() []db.Room {
	return repository.GetRooms()
}

func GetRoomByPlaceId(placeId uint) db.Room {
	return repository.GetRoomByPlaceId(placeId)
}

type RoomStats struct {
	Number string
	IsMale bool
	AreaSqMeters float32
	OccupiedPlaces []db.Place
	FreePlaces []db.Place
	StudentsLiving []db.Student
}

func GetRoomStatsByNumber(number string, room_stats *RoomStats) error {
	var room db.Room
	db.DB.Where("number = ?", number).First(&room)
	if room.Id == 0 {
		return fmt.Errorf("Room with number %s not found", number)
	}
	room_stats.Number = room.Number
	room_stats.IsMale = room.IsMale
	room_stats.AreaSqMeters = room.AreaSqMeters

	var occupiedPlaces []db.Place
	db.DB.Where("room_id = ? AND is_free = ?", room.Id, false).Find(&occupiedPlaces)
	room_stats.OccupiedPlaces = occupiedPlaces

	var freePlaces []db.Place
	db.DB.Where("room_id = ? AND is_free = ?", room.Id, true).Find(&freePlaces)
	room_stats.FreePlaces = freePlaces

	var occupiedPlaceIds []uint
	for _, place := range occupiedPlaces {
		occupiedPlaceIds = append(occupiedPlaceIds, place.Id)
	}

	var studentsLiving []db.Student
	db.DB.Where("place_id IN (?)", occupiedPlaceIds).Find(&studentsLiving)
	room_stats.StudentsLiving = studentsLiving

	return nil
}