package service

import (
	"fmt"

	"github.com/it-02/dormitory/repository"
	"github.com/it-02/dormitory/db"
)

type DormitoryLoad struct {
	TotalPlacesAmount int
	FreeMalePlaces []db.Place
	FreeFemalePlaces []db.Place
	OccupiedMalePlaces []db.Place
	OccupiedFemalePlaces []db.Place
}

func GetDormitoryLoad(uuid string) (DormitoryLoad, error) {
	if !IsUserAdmin(uuid) {
		return DormitoryLoad{}, fmt.Errorf("User is not admin")
	}

	var dormitoryLoad DormitoryLoad
	dormitoryLoad.TotalPlacesAmount = len(repository.GetPlaces())
	dormitoryLoad.FreeMalePlaces = repository.GetPlacesByParams(true, true)
	dormitoryLoad.FreeFemalePlaces = repository.GetPlacesByParams(false, true)
	dormitoryLoad.OccupiedMalePlaces = repository.GetPlacesByParams(true, false)
	dormitoryLoad.OccupiedFemalePlaces = repository.GetPlacesByParams(false, false)
	return dormitoryLoad, nil
}
