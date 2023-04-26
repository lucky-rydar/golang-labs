package structs

import (
	"github.com/it-02/dormitory/internals/db"
	"github.com/it-02/dormitory/internals/features/dorm_stats/structs"
)

type StudentRepr struct {
	Id          uint
	Name        string
	Surname     string
	IsMale	    bool
	Place       structs.PlaceRepr
	Contract    db.Contract
	StudentTicket db.StudentTicket
}
