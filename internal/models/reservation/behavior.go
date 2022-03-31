package reservation

import (
	"time"

	"github.com/youssef-aly1996/bookings/internal/models/room"
)

type Reservationer interface {
	AllReservations() ([]Reservation, error)
	NewReservations() ([]Reservation, error)
	UpdateReservation(u Reservation) error	
	DeleteReservation(id int) error
	GetReservationsById(id int) (Reservation, error)
	InsertReservation(res Reservation) (int, error)
	CheckAvailability(start, end time.Time) ([]room.Room, error)
	CheckAvailabilityByRoomId(start, end time.Time, id string) ([]room.Room, error)
}

type Service struct {
	r Reservationer
}

func New(r Reservationer) Service {
	return Service{r: r}
}
