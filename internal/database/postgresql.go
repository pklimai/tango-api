package database

import (
	"context"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib" // nolint:revive
	"github.com/jmoiron/sqlx"
	"gitlab.com/zigal0-group/nica/tango-api/config"
	"gitlab.com/zigal0/architect/pkg/closer"
)

func InitPostgreSQLConn(ctx context.Context) (*sqlx.DB, error) {
	pgDSN, err := config.GetValue(config.Key_PGDSN)
	if err != nil {
		return nil, err
	}

	postgresDB, err := sqlx.ConnectContext(ctx, "pgx", pgDSN.String())
	if err != nil {
		return nil, fmt.Errorf("sqlx.ConnectContext: %w", err)
	}

	closer.Add(postgresDB.Close)

	err = postgresDB.Ping()
	if err != nil {
		return nil, fmt.Errorf("postgresDB.Ping: %w", err)
	}

	return postgresDB, nil
}
