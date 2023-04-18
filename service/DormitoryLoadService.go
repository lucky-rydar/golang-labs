package service

import (
	"fmt"

	"github.com/it-02/dormitory/db"
	"github.com/it-02/dormitory/repository"
)

type DormitoryLoadService struct {
	place_repository repository.IPlace
	room_repository IRoom
	user_service IUserService
}

func NewDormitoryLoadService(place_repository repository.IPlace, room_repository IRoom, user_service IUserService) *DormitoryLoadService {
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

func (dms *DormitoryLoadService) GetDormitoryLoad(uuid string) (DormitoryLoad, error) {
	if !dms.user_service.IsUserAdmin(uuid) {
		return DormitoryLoad{}, fmt.Errorf("User is not admin")
	}

	var dormitoryLoad DormitoryLoad
	dormitoryLoad.TotalPlacesAmount = len(dms.place_repository.GetPlaces())
	freeMalePlaces := dms.place_repository.GetPlacesByParams(true, true)
	for i := 0; i < len(freeMalePlaces); i++ {
		place := freeMalePlaces[i]
		room := db.Room{}
		err := dms.room_repository.GetRoomById(place.RoomId, &room)
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

	freeFemalePlaces := dms.place_repository.GetPlacesByParams(false, true)
	for i := 0; i < len(freeFemalePlaces); i++ {
		place := freeFemalePlaces[i]
		room := db.Room{}
		err := dms.room_repository.GetRoomById(place.RoomId, &room)
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

	occupiedMalePlaces := dms.place_repository.GetPlacesByParams(true, false)
	for i := 0; i < len(occupiedMalePlaces); i++ {
		place := occupiedMalePlaces[i]
		room := db.Room{}
		err := dms.room_repository.GetRoomById(place.RoomId, &room)
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

	occupiedFemalePlaces := dms.place_repository.GetPlacesByParams(false, false)
	for i := 0; i < len(occupiedFemalePlaces); i++ {
		place := occupiedFemalePlaces[i]
		room := db.Room{}
		err := dms.room_repository.GetRoomById(place.RoomId, &room)
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
