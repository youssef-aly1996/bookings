package roomrestriction

import (
	"time"

	"github.com/youssef-aly1996/bookings/internal/models/room"
)

type RoomRestrictioner interface {
	InsertRoomRestriction(r RoomRestriction) error
	SearchRoomRestriction(start, end time.Time) ([]room.Room, error)
}

type Service struct {
	r RoomRestrictioner
}

func New(r RoomRestrictioner) Service {
	return Service{r: r}
}
