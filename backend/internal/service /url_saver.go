package service_

import (
	"fmt"
	converter "github.com/goriiin/myapp/backend/internal/converter/url"
	"github.com/goriiin/myapp/backend/internal/repository/postgres"
	"github.com/goriiin/myapp/backend/pkg/random"
)

type Url struct {
	Url   string `json:"url" validate:"required,url"`
	Alias string `json:"alias,omitempty"`
}

type urlSaver struct {
	storage URLRepository
}

// URLRepository TODO: структуру postgres.URL в другое место
type URLRepository interface {
	SaveURL(urlToSave string, alias string) error
	RemoveURL(urlToRemove string) error
	EditURL(savedURL string, newAlias string) error
	GetURL(alias string) (*postgres.URL, error)
}

func NewUrlSaverService(storage URLRepository) *urlSaver {
	return &urlSaver{
		storage: storage,
	}
}

func (u *urlSaver) SaveURL(urlToSave string, alias string) error {

	const op = "service.Saver.SaveURL"
	if urlToSave == "" {
		return fmt.Errorf("op: %s - empty url to save", op)
	}

	if alias == "" {
		alias = random.GetUniqueAlias(urlToSave)
	}

	err := u.storage.SaveURL(urlToSave, alias)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (u *urlSaver) RemoveURL(urlToRemove string) error {
	const op = "service.Shortener.RemoveURL"

	err := u.storage.RemoveURL(urlToRemove)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

// EditURL TODO: возврат измененного Alias
func (u *urlSaver) EditURL(savedURL string, newAlias string) error {
	const op = "service.Shortener.EditURL"
	if newAlias == "" {
		return fmt.Errorf("op: %s - empty new alias", op)
	}
	if savedURL == "" {
		return fmt.Errorf("op: %s - empty url", op)
	}

	return u.storage.EditURL(savedURL, newAlias)
}

func (u *urlSaver) GetURL(alias string) (*Url, error) {
	const op = "service.shortener.urlSaver.GetURL"

	storeURL, err := u.storage.GetURL(alias)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return converter.StorageToService(storeURL), nil
}
