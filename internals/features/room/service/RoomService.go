package service

import (
	"fmt"

	"github.com/it-02/dormitory/internals/db"
	"github.com/it-02/dormitory/internals/features/room/structs"
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

type IPlace interface {
	GetPlaces() []db.Place
	GetFreePlaces() []db.Place
	GetFreePlacesByRoomId(roomId uint) []db.Place
	GetOccupiedPlacesByRoomId(roomId uint) []db.Place
	GetPlacesByRoomId(roomId uint) []db.Place
	GetPlaceById(id uint, place *db.Place) error
	GetPlacesByParams(isMale bool, isFree bool) []db.Place
}

type IStudent interface {
	AddStudent(student *db.Student) error
	SetContract(student_id uint, contract_id uint) error
	GetStudents() []db.Student
	GetStudentById(id uint) db.Student
	GetStudentByTicketId(ticket_id uint) db.Student
	SetStudentToPlace(student_id uint, place_id uint) error
	UnsetStudentFromPlace(student_id uint) error
	GetStudentsByPlaceIds(place_ids []uint) []db.Student
}

type IUserService interface {
	IsUserAdmin(uuid string) bool
}

type RoomService struct {
	room_repository IRoom
	place_repository IPlace
	student_repository IStudent
	user_service IUserService
}

func NewRoomService(room_repository IRoom, place_repository IPlace, student_repository IStudent, user_service IUserService) *RoomService {
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

func (rs *RoomService) GetRoomStatsByNumber(number string, room_stats *structs.RoomStats) error {
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