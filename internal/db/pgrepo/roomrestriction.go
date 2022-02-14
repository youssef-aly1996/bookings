package pgrepo

import (
	"context"
	"time"

	"github.com/youssef-aly1996/bookings/internal/models/room"
	"github.com/youssef-aly1996/bookings/internal/models/roomrestriction"
)

const (
	insertRR = `insert into room_restrictions (start_date, end_date, 
		room_id, reservation_id, created_at, updated_at) values 
		($1,$2,$3,$4,$5,$6)`
	search = `select id, room_name
			  from rooms
	          where id not in 
	(select id from reservations where $1 < end_date and $2 > start_date)`
)

func (pgr *PgRepo) InsertRoomRestriction(r roomrestriction.RoomRestriction) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	_, err := pgr.DbPool.Exec(ctx, insertRR,
		r.StartDate,
		r.EndDate,
		r.RoomId,
		r.ReservationId,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		return err
	}
	return nil
}

func (pgr *PgRepo) SearchRoomRestriction(start, end time.Time) ([]room.Room, error) {
	var rooms []room.Room
	var room room.Room
	rows, err := pgr.DbPool.Query(ctx, search, start, end)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&room.Id, &room.RoomName)
		if err != nil {
			return nil, err
		}
		rooms = append(rooms, room)
	}
	return rooms, nil
}
