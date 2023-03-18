package logic

import (
	"github.com/it-02/dormitory/models"
)

type DormitoryLoad struct {
	TotalPlacesAmount int
	FreeMalePlaces []models.Place
	FreeFemalePlaces []models.Place
	OccupiedMalePlaces []models.Place
	OccupiedFemalePlaces []models.Place
}

func GetDormitoryLoad() DormitoryLoad {
	var dormitoryLoad DormitoryLoad
	dormitoryLoad.TotalPlacesAmount = len(GetPlaces())
	dormitoryLoad.FreeMalePlaces = GetPlacesByParams(true, true)
	dormitoryLoad.FreeFemalePlaces = GetPlacesByParams(false, true)
	dormitoryLoad.OccupiedMalePlaces = GetPlacesByParams(true, false)
	dormitoryLoad.OccupiedFemalePlaces = GetPlacesByParams(false, false)
	return dormitoryLoad
}
