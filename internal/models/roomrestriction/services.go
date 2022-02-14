package roomrestriction

import (
	"time"

	"github.com/youssef-aly1996/bookings/internal/models/room"
)

func (rr Service) Insert(r RoomRestriction) error {
	err := rr.r.InsertRoomRestriction(r)
	if err != nil {
		return err
	}

	return nil
}

func (rr Service) Search(start, end time.Time) ([]room.Room, error) {
	rooms, err := rr.r.SearchRoomRestriction(start, end)
	if err != nil {
		return nil, err
	}
	return rooms, nil
}
