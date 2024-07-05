package postgres

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"time"
)

func (s *Storage) getUniqueAlias(str string) string {
	const op = "storage.postgres.getUniqueAlias"
	const query = `
		select id from url
		where alias = $1
		limit 1;
	`
	var err error
	for {
		hash := sha256.Sum256([]byte(str))
		shorter := base64.RawURLEncoding.EncodeToString(hash[:6])

		row := s.db.QueryRow(context.Background(), query, string(shorter))

		_ = row.Scan(&err)
		if err == nil {
			return shorter
		}

		str += "1"
	}
}

func (s *Storage) SaveURL(urlToSave string) (int64, error) {
	const op = "storage.postgres.SaveURL"

	alias := s.getUniqueAlias(urlToSave)
	_, err := s.db.Exec(context.Background(),
		`insert into url(url, alias, created_at) 
			values ($1, $2, $3)`,
		urlToSave, alias, time.Now())
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	return 0, nil
}

func (s *Storage) SaveURLWithAlias(urlToSave string, alias string) (int64, error) {
	const op = "storage.postgres.SaveURLWithAlias"

	return 0, nil
}

func (s *Storage) RemoveURL(urlToSave string) {
	const op = "storage.postgres.RemoveURL"
}

func (s *Storage) EditURL(savedURL string) error {
	const op = "storage.postgres.EditURL"

	return nil
}

func (s *Storage) GetURL(alias string) error {
	const op = "storage.postgres.GetURL"

	return nil
}

func (s *Storage) GetAlias(savedURL string) error {
	const op = "storage.postgres.GetAlias"

	return nil
}
