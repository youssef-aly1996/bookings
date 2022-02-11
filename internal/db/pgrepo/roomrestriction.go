package pgrepo

import "github.com/youssef-aly1996/bookings/internal/models/roomrestriction"

const (
	insertRR = `insert into room_restrictions (start_date, end_date, 
		room_id, reservation_id, created_at, updated_at) values 
		($1,$2,$3,$4,$5,$6)`
)

func (pgr *PgRepo) InsertRoomRestriction(r roomrestriction.RoomRestriction) error {
	_, err := pgr.DbPool.Exec(ctx, insertRR)
	if err != nil {
		return err
	}
	return nil
}
