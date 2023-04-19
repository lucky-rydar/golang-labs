package repository

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
