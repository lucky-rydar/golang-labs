package repository

import (
	"fmt"

	"github.com/it-02/dormitory/internals/db"
	"gorm.io/gorm"
)

type Room struct {
	db *gorm.DB
}

func NewRoom(db *gorm.DB) *Room {
	return &Room{db: db}
}

func (r *Room) AddRoom(room *db.Room) {
	r.db.Create(room)
	fmt.Printf("Room {%d, %t, %f} inserted\n", room.Id, room.IsMale, room.AreaSqMeters)

	placesCount := int(room.AreaSqMeters / 4)
	for i := 0; i < placesCount; i++ {
		place := db.Place{
			RoomId: room.Id,
			IsFree: true,
		}
		r.db.Create(&place)
	}
}

func (r *Room) GetRooms() []db.Room {
	var rooms []db.Room
	r.db.Find(&rooms)
	return rooms
}

func (r *Room) GetRoomByPlaceId(placeId uint) db.Room {
	var place db.Place
	r.db.First(&place, placeId)
	var room db.Room
	r.db.First(&room, place.RoomId)
	return room
}

func (r *Room) GetRoomById(id uint, room *db.Room) error {
	var err error
	r.db.First(&room, id)
	if room.Id == 0 {
		err = fmt.Errorf("Room with id %d not found", id)
	}
	return err
}

func (r *Room) IsRoomNumberExists(number string) bool {
	var room db.Room
	r.db.Where("number = ?", number).First(&room)
	return room.Id != 0
}

func (r *Room) RemoveRoomById(id uint) error {
	var err error
	var room db.Room
	err = r.GetRoomById(id, &room)
	if err == nil {
		r.db.Delete(&room)
	}
	return err
}

func (r *Room) GetRoomByNumber(number string) db.Room {
	var room db.Room
	r.db.Where("number = ?", number).First(&room)
	return room
}
