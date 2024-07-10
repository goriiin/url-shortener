package repository

import "time"

type URL struct {
	Id        int64     `db:"id"`
	Url       string    `db:"url"`
	CreatedAt time.Time `db:"created_at"`
	Alias     string    `db:"alias"`
}

var URLRepository interface {
	SaveURL(urlToSave string, alias string) error
	RemoveURL(urlToRemove string) error
	EditURL(savedURL string, newAlias string) error
	GetURL(alias string) (*URL, error)
}
