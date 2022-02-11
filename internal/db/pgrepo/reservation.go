package pgrepo

import (
	"context"
	"time"

	"github.com/youssef-aly1996/bookings/internal/models/reservation"
)

const (
	insert = `insert into reservations (first_name, last_name, email
		, phone, start_date, end_date, room_id, created_at, updated_at) values 
		($1,$2,$3,$4,$5,$6,$7,$8,$9) returning id`
)

func (pgr *PgRepo) InsertReservation(res reservation.Reservation) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	var id int
	err := pgr.DbPool.QueryRow(ctx,
		insert,
		res.FirstName,
		res.LastName,
		res.Email,
		res.Phone,
		res.StartDate,
		res.EndDate,
		res.RoomId,
		time.Now(),
		time.Now(),
	).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
