package postgresdriver

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	maxOpenDbConn = 10
	maxIdleDbConn = 5
	maxDbLifetime = 5 * time.Minute
)

var (
	ctx = context.Background()
)

func connect() (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig("host=localhost port=5432 dbname=bookings user=postgres password=joe1234 sslmode=mode")
	config.MaxConns = maxOpenDbConn
	config.MaxConnIdleTime = maxIdleDbConn
	config.MaxConnLifetime = maxDbLifetime
	if err != nil {
		return nil, err
	}
	fmt.Println(*config)
	dbPool, err := pgxpool.ConnectConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	if err = dbPool.Ping(ctx); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return dbPool, nil
}
