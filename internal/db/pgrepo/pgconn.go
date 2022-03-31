package pgrepo

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/spf13/viper"
)

type PgRepo struct {
	DbPool *pgxpool.Pool
}

const (
	// connStr         = "host=localhost port=5432 user=postgres password=joe1234 dbname=bookings"
	maxConns        = 10
	maxConnIdleTime = 5 * time.Second
	MaxConnLifetime = 5 * time.Second
)

var (
	ctx = context.Background()
)

func NewPgRepo() (*PgRepo, error) {
	db, err := newPgConn()
	if err != nil {
		return nil, err
	}
	return &PgRepo{DbPool: db}, nil
}

func newPgConn() (*pgxpool.Pool, error) {
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()

	if err != nil {
	  log.Fatalf("Error while reading config file %s", err)
	}
	con := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", 
	viper.GetString("dbHost"), viper.GetString("dbPort"), viper.GetString("dbUser"), 
	viper.GetString("dbPass"), viper.GetString("dbName"), viper.GetString("dbSSl"))
	fmt.Println(con)
	config, err := pgxpool.ParseConfig(con)
	if err != nil {
		return nil, err
	}
	config.MaxConns = maxConns
	config.MaxConnIdleTime = maxConnIdleTime
	config.MaxConnLifetime = MaxConnLifetime

	dbPool, err := pgxpool.ConnectConfig(ctx, config)
	if err != nil {
		return nil, err
	}
	if err = dbPool.Ping(ctx); err != nil {
		return nil, err
	}
	return dbPool, nil
}
