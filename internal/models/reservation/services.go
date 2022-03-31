package reservation

import (
	"time"

	"github.com/youssef-aly1996/bookings/internal/models/room"
)

func (s Service) Insert(res Reservation) (int, error) {
	id, err := s.r.InsertReservation(res)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s Service) Check(start, end time.Time) ([]room.Room, error) {
	rooms, err := s.r.CheckAvailability(start, end)
	if err != nil {
		return nil, err
	}
	return rooms, nil
}

func (s Service) CheckByRoomId(start, end time.Time, id string) ([]room.Room, error) {
	rooms, err := s.r.CheckAvailabilityByRoomId(start, end, id)
	if err != nil {
		return nil, err
	}
	return rooms, nil
}

func (s Service) All() ([]Reservation, error) {
	ress, err := s.r.AllReservations()
	if err != nil {
		return nil, err
	}
	return ress, nil
}

func (s Service) New() ([]Reservation, error) {
	ress,  err := s.r.NewReservations()
	if err != nil {
		return nil, err
	}
	return ress, nil
}

func (s Service) GetByID(id int) (Reservation, error) {
	ress,  err := s.r.GetReservationsById(id)
	if err != nil {
		return ress, err
	}
	return ress, nil
}

func (s Service) Update(u Reservation) error {
	err := s.r.UpdateReservation(u)
	if err != nil {
		return err
	}
	return nil
}

func (s Service) Delete(id int) error {
	err := s.r.DeleteReservation(id)
	if err != nil {
		return err
	}
	return nil
}