package repository

import (
	"fmt"

	"github.com/it-02/dormitory/db"
)

func AddRoom(room db.Room) {
	db.DB.Create(&room)
	fmt.Printf("Room {%d, %t, %f} inserted\n", room.Id, room.IsMale, room.AreaSqMeters)

	placesCount := int(room.AreaSqMeters / 4)
	for i := 0; i < placesCount; i++ {
		place := db.Place{
			RoomId: room.Id,
			IsFree: true,
		}
		db.DB.Create(&place)
	}
}

func GetRooms() []db.Room {
	var rooms []db.Room
	db.DB.Find(&rooms)
	return rooms
}

func GetRoomByPlaceId(placeId uint) db.Room {
	var place db.Place
	db.DB.First(&place, placeId)
	var room db.Room
	db.DB.First(&room, place.RoomId)
	return room
}

func GetRoomById(id uint, room *db.Room) error {
	var err error
	db.DB.First(&room, id)
	if room.Id == 0 {
		err = fmt.Errorf("Room with id %d not found", id)
	}
	return err
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
