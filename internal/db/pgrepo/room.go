package pgrepo

const (
	get = `select room_name from rooms where id=$1`
)

func (pgr *PgRepo) GetRoomById(id int) (string, error) {
	var roomName string
	row := pgr.DbPool.QueryRow(ctx, get, id)
	err := row.Scan(&roomName)
	if err != nil {
		return "", err
	}
	return roomName, nil
}
