package logic

import (
	"fmt"

	"github.com/it-02/dormitory/db"
	"github.com/it-02/dormitory/models"
)

func AddRoom(room models.Room) {
	db.DB.Create(&room)
	fmt.Printf("Room {%d, %t, %f} inserted\n", room.Id, room.IsMale, room.AreaSqMeters)
}

func GetRooms() []models.Room {
	var rooms []models.Room
	db.DB.Find(&rooms)
	return rooms
}

func GetRoomByPlaceId(placeId uint) models.Room {
	var place models.Place
	db.DB.First(&place, placeId)
	var room models.Room
	db.DB.First(&room, place.RoomId)
	return room
}

func GetRoomById(id uint, room *models.Room) error {
	var err error
	db.DB.First(&room, id)
	if room.Id == 0 {
		err = fmt.Errorf("Room with id %d not found", id)
	}
	return err
}
