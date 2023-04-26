package repository

import (
	"fmt"

	"github.com/it-02/dormitory/internals/db"
	"gorm.io/gorm"
)

type IPlace interface {
	GetPlaces() []db.Place
	GetFreePlaces() []db.Place
	GetFreePlacesByRoomId(roomId uint) []db.Place
	GetOccupiedPlacesByRoomId(roomId uint) []db.Place
	GetPlacesByRoomId(roomId uint) []db.Place
	GetPlaceById(id uint, place *db.Place) error
	GetPlacesByParams(isMale bool, isFree bool) []db.Place
}

type Place struct {
	db *gorm.DB
}

func NewPlace(db *gorm.DB) *Place {
	return &Place{db: db}
}

func (p *Place) GetPlaces() []db.Place {
	var places []db.Place
	p.db.Find(&places)
	return places
}

func (p *Place) GetFreePlaces() []db.Place {
	var places []db.Place
	p.db.Where("is_free = ?", true).Find(&places)
	return places
}

func (p *Place) GetFreePlacesByRoomId(roomId uint) []db.Place {
	var places []db.Place
	p.db.Where("is_free = ? AND room_id = ?", true, roomId).Find(&places)
	return places
}

func (p *Place) GetOccupiedPlacesByRoomId(roomId uint) []db.Place {
	var places []db.Place
	p.db.Where("is_free = ? AND room_id = ?", false, roomId).Find(&places)
	return places
}

func (p *Place) GetPlacesByRoomId(roomId uint) []db.Place {
	var places []db.Place
	p.db.Where("room_id = ?", roomId).Find(&places)
	return places
}

func (p *Place) GetPlaceById(id uint, place *db.Place) error {
	var err error
	p.db.First(&place, id)
	if place.Id == 0 {
		err = fmt.Errorf("Place with id %d not found", id)
	}
	return err
}

func (p *Place) GetPlacesByParams(isMale bool, isFree bool) []db.Place {
	var places []db.Place
	p.db.Where("is_free = ?", isFree).Find(&places)

	var placesRet []db.Place
	for i := 0; i < len(places); i++ {
		var room db.Room
		p.db.First(&room, places[i].RoomId)
		if room.IsMale == isMale {
			placesRet = append(placesRet, places[i])
		}
	}

	return placesRet
}
