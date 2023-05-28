package service

import (
	"testing"

	"github.com/it-02/dormitory/internals/db"
	
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RoomMock struct {
	rooms []db.Room
}

func (rm *RoomMock) AddRoom(room *db.Room) {
	rm.rooms = append(rm.rooms, *room)
}

func (rm *RoomMock) GetRooms() []db.Room {
	return rm.rooms
}

func (rm *RoomMock) GetRoomByPlaceId(placeId uint) db.Room {
	for _, room := range rm.rooms {
		for _, place := range room.Places {
			if place.Id == placeId {
				return room
			}
		}
	}
	return db.Room{}
}

func (rm *RoomMock) IsRoomNumberExists(roomNumber string) bool {
	for _, room := range rm.rooms {
		if room.Number == roomNumber {
			return true
		}
	}
	return false
}

func (rm *RoomMock) GetRoomByNumber(number string) db.Room {
	for _, room := range rm.rooms {
		if room.Number == number {
			return room
		}
	}
	return db.Room{}
}

func NewRoomMock() *RoomMock {
	return &RoomMock{rooms: []db.Room{}}
}

/*
type IPlace interface {
	GetPlaces() []db.Place
	GetFreePlaces() []db.Place
	GetFreePlacesByRoomId(roomId uint) []db.Place
	GetOccupiedPlacesByRoomId(roomId uint) []db.Place
}
*/
type PlaceMock struct {
	places []db.Place


func NewUserServiceMock() *UserServiceMock {
	return &UserServiceMock{}
}

type RoomServiceTestSuite struct {
	suite.Suite
	room_service RoomService
}

func (suite *RoomServiceTestSuite) SetupTest() {
	room_service := NewRoomService(NewRoomMock(), NewUserServiceMock())
}



func TestRoomService(t *testing.T) {
	suite.Run(t, new(RoomServiceTestSuite))
}
