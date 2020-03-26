package postgres

import (
	"context"
	"github.com/aibotsoft/micro/config"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
)

func New(cfg *config.Config) (*pgx.Conn, error) {
	ctx, cancel := context.WithTimeout(context.Background(), cfg.Postgres.Timeout)
	defer cancel()

	connConfig, err := pgx.ParseConfig("")
	if err != nil {
		return nil, errors.Wrap(err, "pgx.ParseConfig error")
	}
	conn, err := pgx.ConnectConfig(ctx, connConfig)
	if err != nil {
		return nil, errors.Wrap(err, "pgx.ConnectConfig error")
	}
	err = conn.Ping(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "postgres ping error")
	}
	return conn, nil
}
