package postgres

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
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

func (s *Storage) getUniqueAlias(str string) string {
	const query = `
		select id from url
		where alias = $1
		limit 1;
	`
	var err error
	for {
		hash := sha256.Sum256([]byte(str))
		shorter := base64.RawURLEncoding.EncodeToString(hash[:6])

		row := s.db.QueryRow(context.Background(), query, shorter)

		_ = row.Scan(&err)
		if err == nil {
			return shorter
		}

		str += "1"
	}
}

func (s *Storage) SaveURL(urlToSave string) error {
	const op = "storage.postgres.SaveURL"

	alias := s.getUniqueAlias(urlToSave)
	_, err := s.db.Exec(context.Background(), insert,
		urlToSave, alias, time.Now())
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Storage) SaveURLWithAlias(urlToSave string, alias string) (int64, error) {
	const op = "storage.postgres.SaveURLWithAlias"
	_, err := s.db.Exec(context.Background(),
		insert, urlToSave, alias, time.Now())
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return 0, nil
}

func (s *Storage) RemoveURL(savedURL string) error {
	const op = "storage.postgres.RemoveURL"

	_, err := s.db.Exec(context.Background(), deleteURL, savedURL)
	if err != nil {
		return fmt.Errorf("%s: %s\n", op, err)
	}

	return nil
}

func (s *Storage) EditURL(savedURL string, newAlias string) error {
	const op = "storage.postgres.EditURL"

	_, err := s.db.Exec(context.Background(), updateAlias, newAlias, savedURL)
	if err != nil {
		return fmt.Errorf("%s: %s\n", op, err)
	}

	return nil
}

func (s *Storage) GetURL(alias string) (*URL, error) {
	const op = "storage.postgres.GetURL"

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
	const op = "storage.postgres.GetAlias"
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
