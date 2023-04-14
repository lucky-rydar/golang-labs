package repository

import (
	"fmt"

	"github.com/it-02/dormitory/db"
	"gorm.io/gorm"
)

type IRoom interface {
	AddRoom(room *db.Room)
	GetRooms() []db.Room
	GetRoomByPlaceId(placeId uint) db.Room
	GetRoomById(id uint, room *db.Room) error
	IsRoomNumberExists(roomNumber string) bool
	RemoveRoomById(id uint) error
	GetRoomByNumber(number string) db.Room
}

type Room struct {
	db *gorm.DB
}

func NewRoom(db *gorm.DB) IRoom {
	return &Room{db: db}
}

func (this Room) AddRoom(room *db.Room) {
	this.db.Create(room)
	fmt.Printf("Room {%d, %t, %f} inserted\n", room.Id, room.IsMale, room.AreaSqMeters)

	placesCount := int(room.AreaSqMeters / 4)
	for i := 0; i < placesCount; i++ {
		place := db.Place{
			RoomId: room.Id,
			IsFree: true,
		}
		this.db.Create(&place)
	}
}

func (this Room) GetRooms() []db.Room {
	var rooms []db.Room
	this.db.Find(&rooms)
	return rooms
}

func (this Room) GetRoomByPlaceId(placeId uint) db.Room {
	var place db.Place
	this.db.First(&place, placeId)
	var room db.Room
	this.db.First(&room, place.RoomId)
	return room
}

func (this Room) GetRoomById(id uint, room *db.Room) error {
	var err error
	this.db.First(&room, id)
	if room.Id == 0 {
		err = fmt.Errorf("Room with id %d not found", id)
	}
	return err
}

func (this Room) IsRoomNumberExists(number string) bool {
	var room db.Room
	this.db.Where("number = ?", number).First(&room)
	return room.Id != 0
}

func (this Room) RemoveRoomById(id uint) error {
	var err error
	var room db.Room
	err = this.GetRoomById(id, &room)
	if err == nil {
		this.db.Delete(&room)
	}
	return err
}

func (this Room) GetRoomByNumber(number string) db.Room {
	var room db.Room
	this.db.Where("number = ?", number).First(&room)
	return room
}
