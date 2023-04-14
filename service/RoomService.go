package service

import (
	"fmt"

	"github.com/it-02/dormitory/repository"
	"github.com/it-02/dormitory/db"
)

type IRoomService interface {
	AddRoom(uuid string, room *db.Room) error
	GetRooms() []db.Room
	GetRoomByPlaceId(placeId uint) db.Room
	GetRoomStatsByNumber(number string, room_stats *RoomStats) error
}

type RoomService struct {
	room_repository repository.IRoom
	place_repository repository.IPlace
	student_repository repository.IStudent
	user_service IUserService
}

func NewRoomService(room_repository repository.IRoom, place_repository repository.IPlace, student_repository repository.IStudent, user_service IUserService) IRoomService {
	return &RoomService{room_repository: room_repository, place_repository: place_repository, student_repository: student_repository, user_service: user_service}
}

func (this RoomService) AddRoom(uuid string, room *db.Room) error {
	if !this.user_service.IsUserAdmin(uuid) {
		return fmt.Errorf("User is not admin")
	}

	if !this.room_repository.IsRoomNumberExists(room.Number) {
		this.room_repository.AddRoom(room)
	} else {
		return fmt.Errorf("Room with number %s already exists", room.Number)
	}
	return nil
}

func (this RoomService) GetRooms() []db.Room {
	return this.room_repository.GetRooms()
}

func (this RoomService) GetRoomByPlaceId(placeId uint) db.Room {
	return this.room_repository.GetRoomByPlaceId(placeId)
}

type RoomStats struct {
	Number string
	IsMale bool
	AreaSqMeters float32
	OccupiedPlaces []db.Place
	FreePlaces []db.Place
	StudentsLiving []db.Student
}

func (this RoomService) GetRoomStatsByNumber(number string, room_stats *RoomStats) error {
	room := this.room_repository.GetRoomByNumber(number)
	if room.Id == 0 {
		return fmt.Errorf("Room with number %s not found", number)
	}

	room_stats.Number = room.Number
	room_stats.IsMale = room.IsMale
	room_stats.AreaSqMeters = room.AreaSqMeters

	occupiedPlaces := this.place_repository.GetOccupiedPlacesByRoomId(room.Id)
	room_stats.OccupiedPlaces = occupiedPlaces

	freePlaces := this.place_repository.GetFreePlacesByRoomId(room.Id)
	room_stats.FreePlaces = freePlaces

	var occupiedPlaceIds []uint
	for _, place := range occupiedPlaces {
		occupiedPlaceIds = append(occupiedPlaceIds, place.Id)
	}

	studentsLiving := this.student_repository.GetStudentsByPlaceIds(occupiedPlaceIds)
	room_stats.StudentsLiving = studentsLiving

	return nil
}