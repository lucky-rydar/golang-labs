package service

import (
	"fmt"

	"github.com/it-02/dormitory/repository"
	"github.com/it-02/dormitory/db"
)

type RoomService struct {
	room_repository repository.IRoom
	place_repository repository.IPlace
	student_repository repository.IStudent
	user_service IUserService
}

func NewRoomService(room_repository repository.IRoom, place_repository repository.IPlace, student_repository repository.IStudent, user_service IUserService) *RoomService {
	return &RoomService{room_repository: room_repository, place_repository: place_repository, student_repository: student_repository, user_service: user_service}
}

func (rs *RoomService) AddRoom(uuid string, room *db.Room) error {
	if !rs.user_service.IsUserAdmin(uuid) {
		return fmt.Errorf("User is not admin")
	}

	if !rs.room_repository.IsRoomNumberExists(room.Number) {
		rs.room_repository.AddRoom(room)
	} else {
		return fmt.Errorf("Room with number %s already exists", room.Number)
	}
	return nil
}

func (rs *RoomService) GetRooms() []db.Room {
	return rs.room_repository.GetRooms()
}

func (rs *RoomService) GetRoomByPlaceId(placeId uint) db.Room {
	return rs.room_repository.GetRoomByPlaceId(placeId)
}

type RoomStats struct {
	Number string
	IsMale bool
	AreaSqMeters float32
	OccupiedPlaces []db.Place
	FreePlaces []db.Place
	StudentsLiving []db.Student
}

func (rs *RoomService) GetRoomStatsByNumber(number string, room_stats *RoomStats) error {
	room := rs.room_repository.GetRoomByNumber(number)
	if room.Id == 0 {
		return fmt.Errorf("Room with number %s not found", number)
	}

	room_stats.Number = room.Number
	room_stats.IsMale = room.IsMale
	room_stats.AreaSqMeters = room.AreaSqMeters

	occupiedPlaces := rs.place_repository.GetOccupiedPlacesByRoomId(room.Id)
	room_stats.OccupiedPlaces = occupiedPlaces

	freePlaces := rs.place_repository.GetFreePlacesByRoomId(room.Id)
	room_stats.FreePlaces = freePlaces

	var occupiedPlaceIds []uint
	for _, place := range occupiedPlaces {
		occupiedPlaceIds = append(occupiedPlaceIds, place.Id)
	}

	studentsLiving := rs.student_repository.GetStudentsByPlaceIds(occupiedPlaceIds)
	room_stats.StudentsLiving = studentsLiving

	return nil
}