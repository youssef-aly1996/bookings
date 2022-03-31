package pgrepo

import (
	"context"
	"time"

	"github.com/youssef-aly1996/bookings/internal/models/reservation"
	"github.com/youssef-aly1996/bookings/internal/models/room"
)

const (
	insert = `insert into reservations (first_name, last_name, email
		, phone, start_date, end_date, room_id, created_at, updated_at) values 
		($1,$2,$3,$4,$5,$6,$7,$8,$9) returning id`
	searchAvailability = `
	select id, room_name
	from rooms
	where room_name not in (select rm.room_name
	from rooms rm
	join reservations rs on rm.id = rs.room_id 
	where $1 between start_date and end_date or $2 between start_date and end_date);`

	searchAvailabilityByRoomId = `select id, room_name
	from rooms
	where room_name not in (select rm.room_name
	from rooms rm
	join reservations rs on rm.id = rs.room_id 
	where $1 between start_date and end_date or $2 between start_date and end_date)
	and id = $3`
	allRes = `select rs.*, room_name 
	from reservations rs
	left join rooms on rs.room_id = rooms.id`
	newRes = `select rs.*, room_name 
	from reservations rs
	left join rooms on rs.room_id = rooms.id
	order by created_at desc`
	resById = `select rs.*, room_name 
	from reservations rs
	left join rooms on rs.room_id = rooms.id
	where rs.id = $1`
	updateRes = `
	update reservations set first_name = $1, last_name = $2, email = $3, phone = $4, updated_at = $5
	where id = $6`	
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

func (pgr *PgRepo) CheckAvailability(start, end time.Time) ([]room.Room, error) {
	var rooms []room.Room
	var room room.Room
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	rows, err := pgr.DbPool.Query(ctx, searchAvailability, start, end)
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

func (pgr *PgRepo) CheckAvailabilityByRoomId(start, end time.Time, id string) ([]room.Room, error) {
	var rooms []room.Room
	var room room.Room
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	rows, err := pgr.DbPool.Query(ctx, searchAvailabilityByRoomId, start, end, id)
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

func (pgr *PgRepo) AllReservations() ([]reservation.Reservation, error) {
	var reservations []reservation.Reservation
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	rows, err := pgr.DbPool.Query(ctx, allRes)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var reservation reservation.Reservation
		err := rows.Scan(
			&reservation.Id, 
			&reservation.FirstName, 
			&reservation.LastName, 
			&reservation.Email,
			&reservation.Phone,
			&reservation.StartDate,
			&reservation.EndDate,
			&reservation.RoomId,	
			&reservation.CreatedAt,
			&reservation.UpdatedAt,
			&reservation.RoomName,			
		)
		if err != nil {
			return nil, err
		}
		reservations = append(reservations, reservation)
	}
	return reservations, nil
}


func (pgr *PgRepo) NewReservations() ([]reservation.Reservation, error) {
	var reservations []reservation.Reservation
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	rows, err := pgr.DbPool.Query(ctx, newRes)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var reservation reservation.Reservation
		err := rows.Scan(
			&reservation.Id, 
			&reservation.FirstName, 
			&reservation.LastName, 
			&reservation.Email,
			&reservation.Phone,
			&reservation.StartDate,
			&reservation.EndDate,
			&reservation.RoomId,	
			&reservation.CreatedAt,
			&reservation.UpdatedAt,
			&reservation.RoomName,			
		)
		if err != nil {
			return nil, err
		}
		reservations = append(reservations, reservation)
	}
	return reservations, nil
}

func (pgr *PgRepo) GetReservationsById(id int) (reservation.Reservation, error) {
	var reservation reservation.Reservation
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	row := pgr.DbPool.QueryRow(ctx, resById, id)

			err := row.Scan(
			&reservation.Id, 
			&reservation.FirstName, 
			&reservation.LastName, 
			&reservation.Email,
			&reservation.Phone,
			&reservation.StartDate,
			&reservation.EndDate,
			&reservation.RoomId,	
			&reservation.CreatedAt,
			&reservation.UpdatedAt,
			&reservation.RoomName,			
		)
		if err != nil {
			return reservation, err
		}

	return reservation, nil
}


func (pgr *PgRepo) UpdateReservation(u reservation.Reservation) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := pgr.DbPool.Exec(ctx, updateRes,
		u.FirstName,
		u.LastName,
		u.Email,
		u.Phone,
		time.Now(),
		u.Id,
	)

	if err != nil {
		return err
	}

	return nil
}

func (pgr *PgRepo) DeleteReservation(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "delete from reservations where id = $1"

	_, err := pgr.DbPool.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
