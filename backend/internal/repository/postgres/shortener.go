package postgres

import (
	"context"
	"fmt"
	"github.com/goriiin/myapp/backend/db/postgres"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

const (
	insert      = `insert into url(url, alias, created_at) values ($1, $2, $3);`
	updateAlias = `update url set alias = $1 where url = $2;`
	deleteURL   = `delete from url where url = $1;`
	getByURL    = `select id, url.url, alias, created_at from url where url = $1;`
	getByAlias  = `select id, url.url, alias, created_at from url where alias = $1;`
)

type URL struct {
	Id        int64     `db:"id"`
	Url       string    `db:"url"`
	CreatedAt time.Time `db:"created_at"`
	Alias     string    `db:"alias"`
}

type Storage struct {
	db *pgxpool.Pool
}

func New() *Storage {
	db, err := postgres.New()
	if err != nil {
		panic(err)
	}

	return &Storage{db}
}

func (s *Storage) SaveURL(urlToSave string, alias string) error {
	const op = "repository.postgres.SaveURLWithAlias"
	_, err := s.db.Exec(context.Background(),
		insert, urlToSave, alias, time.Now())
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Storage) RemoveURL(urlToRemove string) error {
	const op = "repository.postgres.RemoveURL"

	_, err := s.db.Exec(context.Background(), deleteURL, urlToRemove)
	if err != nil {
		return fmt.Errorf("%s: %s\n", op, err)
	}

	return nil
}

func (s *Storage) EditURL(savedURL string, newAlias string) error {
	const op = "repository.postgres.EditURL"

	_, err := s.db.Exec(context.Background(), updateAlias, newAlias, savedURL)
	if err != nil {
		return fmt.Errorf("%s: %s\n", op, err)
	}

	return nil
}

func (s *Storage) GetURL(alias string) (*URL, error) {
	const op = "repository.postgres.GetURL"

	row, err := s.db.Query(context.Background(), getByAlias, alias)
	if err != nil {
		return nil, fmt.Errorf("%s: %s\n", op, err)
	}
	row.Close()

	var url URL
	err = row.Scan(&url.Id, &url.Url, &url.Alias, &url.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("%s: %s\n", op, err)
	}

	return &url, nil
}

func (s *Storage) GetAlias(savedURL string) (*URL, error) {
	const op = "repository.postgres.GetAlias"
	row, err := s.db.Query(context.Background(), getByURL, savedURL)
	if err != nil {
		return nil, fmt.Errorf("%s: %s\n", op, err)
	}
	defer row.Close()
	var url URL

	err = row.Scan(&url.Id, &url.Url, &url.Alias, &url.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("%s: %s\n", op, err)
	}

	return &url, nil
}
