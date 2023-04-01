package repository

import (
	"fmt"

	"github.com/it-02/dormitory/db"
)

func AddRoom(room *db.Room) {
	db.DB.Create(room)
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

func IsRoomNumberExists(number string) bool {
	var room db.Room
	db.DB.Where("number = ?", number).First(&room)
	return room.Id != 0
}

func RemoveRoomById(id uint) error {
	var err error
	var room db.Room
	err = GetRoomById(id, &room)
	if err == nil {
		db.DB.Delete(&room)
	}
	return err
}

func GetRoomByNumber(number string) db.Room {
	var room db.Room
	db.DB.Where("number = ?", number).First(&room)
	return room
}
