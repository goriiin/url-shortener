package postgres

import (
	"context"
	"fmt"
	"github.com/goriiin/myapp/backend/db/postgres/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func New() (*pgxpool.Pool, error) {
	const op = "repository.postgres.New"

	cfg, err := config.NewStorageConfig()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	poolConfig, err := config.NewPoolConfig(&cfg)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	poolConfig.MaxConns = cfg.MaxConns

	db, err := config.NewConnection(poolConfig)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	_, err = db.Exec(context.Background(), setDB)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return db, nil
}

const (
	setDB = `
	create table if not exists url 
	(
	    id serial primary key,
	    alias text NOT NULL unique,
	    url text NOT NULL UNIQUE,
	    created_at timestamp not null
	);
`

	rm = `drop table if exists url;`
)

func RmTables(db *pgxpool.Pool) error {
	const op = "repository.postgres.RmTables"
	if _, err := db.Exec(context.Background(), rm); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
