package postgresdriver

import (
	"github.com/jackc/pgx/v4/pgxpool"
)

type PostgresDbRepo struct {
	DbPool *pgxpool.Pool
}

//return concrete type as this is the producer
func NewPostgresRepo() *PostgresDbRepo {
	db, err := connect()
	if err != nil {
		return nil
	}
	return &PostgresDbRepo{DbPool: db}
}

func (pg *PostgresDbRepo) AllUsers() bool {
	return true
}
