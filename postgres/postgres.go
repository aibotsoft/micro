package postgres

import (
	"context"
	"github.com/aibotsoft/micro/config"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
)

// Connect creates a new Pool and immediately establishes one connection.
// Error panics
func MustConnect(cfg *config.Config) *pgxpool.Pool {
	db, err := Connect(cfg)
	if err != nil {
		panic(err)
	}
	return db
}

// Connect creates a new Pool and immediately establishes one connection
func Connect(cfg *config.Config) (*pgxpool.Pool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), cfg.Postgres.Timeout)
	defer cancel()

	connConfig, err := pgxpool.ParseConfig("")
	if err != nil {
		return nil, errors.Wrap(err, "pgxpool.ParseConfig error")
	}
	conn, err := pgxpool.ConnectConfig(ctx, connConfig)
	if err != nil {
		return nil, errors.Wrap(err, "pgxpool.ConnectConfig error")
	}
	return conn, nil
}
