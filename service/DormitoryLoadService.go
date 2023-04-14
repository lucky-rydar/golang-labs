package service

import (
	"fmt"

	"github.com/it-02/dormitory/repository"
	"github.com/it-02/dormitory/db"
)

type IDormitoryLoadService interface {
	GetDormitoryLoad(uuid string) (DormitoryLoad, error)
}

type DormitoryLoadService struct {
	place_repository *repository.IPlace
	room_repository *repository.IRoom
	user_service *IUserService
}

func NewDormitoryLoadService(place_repository *repository.IPlace, room_repository *repository.IRoom, user_service *IUserService) IDormitoryLoad {
	return &DormitoryLoadService{place_repository: place_repository, room_repository: room_repository, user_service: user_service}
}

type PlaceRepr struct {
	PlaceId uint
	IsMale bool
	IsFree bool
	RoomNumber string
}

type DormitoryLoad struct {
	TotalPlacesAmount int
	FreeMalePlaces []PlaceRepr
	FreeFemalePlaces []PlaceRepr
	OccupiedMalePlaces []PlaceRepr
	OccupiedFemalePlaces []PlaceRepr
}

func (this DormitoryLoadService) GetDormitoryLoad(uuid string) (DormitoryLoad, error) {
	if !IsUserAdmin(uuid) {
		return DormitoryLoad{}, fmt.Errorf("User is not admin")
	}

	var dormitoryLoad DormitoryLoad
	dormitoryLoad.TotalPlacesAmount = len(this.place_repository.GetPlaces())
	freeMalePlaces := this.place_repository.GetPlacesByParams(true, true)
	for i := 0; i < len(freeMalePlaces); i++ {
		place := freeMalePlaces[i]
		room := db.Room{}
		err := this.room_repository.GetRoomById(place.RoomId, &room)
		if err != nil {
			return DormitoryLoad{}, err
		}

		place_repr := PlaceRepr{
			PlaceId: place.Id,
			IsFree: place.IsFree,
			IsMale: room.IsMale,
			RoomNumber: room.Number,
		}
		dormitoryLoad.FreeMalePlaces = append(dormitoryLoad.FreeMalePlaces, place_repr)
	}

	freeFemalePlaces := this.place_repository.GetPlacesByParams(false, true)
	for i := 0; i < len(freeFemalePlaces); i++ {
		place := freeFemalePlaces[i]
		room := db.Room{}
		err := this.room_repository.GetRoomById(place.RoomId, &room)
		if err != nil {
			return DormitoryLoad{}, err
		}

		place_repr := PlaceRepr{
			PlaceId: place.Id,
			IsFree: place.IsFree,
			IsMale: room.IsMale,
			RoomNumber: room.Number,
		}
		dormitoryLoad.FreeFemalePlaces = append(dormitoryLoad.FreeFemalePlaces, place_repr)
	}

	occupiedMalePlaces := this.place_repository.GetPlacesByParams(true, false)
	for i := 0; i < len(occupiedMalePlaces); i++ {
		place := occupiedMalePlaces[i]
		room := db.Room{}
		err := this.room_repository.GetRoomById(place.RoomId, &room)
		if err != nil {
			return DormitoryLoad{}, err
		}

		place_repr := PlaceRepr{
			PlaceId: place.Id,
			IsFree: place.IsFree,
			IsMale: room.IsMale,
			RoomNumber: room.Number,
		}
		dormitoryLoad.OccupiedMalePlaces = append(dormitoryLoad.OccupiedMalePlaces, place_repr)
	}

	occupiedFemalePlaces := this.place_repository.GetPlacesByParams(false, false)
	for i := 0; i < len(occupiedFemalePlaces); i++ {
		place := occupiedFemalePlaces[i]
		room := db.Room{}
		err := this.room_repository.GetRoomById(place.RoomId, &room)
		if err != nil {
			return DormitoryLoad{}, err
		}

		place_repr := PlaceRepr{
			PlaceId: place.Id,
			IsFree: place.IsFree,
			IsMale: room.IsMale,
			RoomNumber: room.Number,
		}
		dormitoryLoad.OccupiedFemalePlaces = append(dormitoryLoad.OccupiedFemalePlaces, place_repr)
	}

	return dormitoryLoad, nil
}
