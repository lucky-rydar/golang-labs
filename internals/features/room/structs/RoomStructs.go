package structs

import (
	"github.com/it-02/dormitory/internals/db"
)

type RoomStats struct {
	Number string
	IsMale bool
	AreaSqMeters float32
	OccupiedPlaces []db.Place
	FreePlaces []db.Place
	StudentsLiving []db.Student
}
