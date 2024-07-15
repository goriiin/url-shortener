package service_

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	converter "github.com/goriiin/myapp/backend/internal/converter/url"
	"github.com/goriiin/myapp/backend/internal/repository/postgres"
)

type URL struct {
	Url   string `json:"url"`
	Alias string `json:"alias"`
}

type urlSaver struct {
	storage postgres.Storage
}

func NewUrlSaverService(storage postgres.Storage) *urlSaver {
	return &urlSaver{
		storage: storage,
	}
}

func getUniqueAlias(str string) string {
	hash := sha256.Sum256([]byte(str))
	return base64.RawURLEncoding.EncodeToString(hash[:6])
}

func (u *urlSaver) SaveURL(urlToSave string, alias string) error {
	const op = "service.Saver.SaveURL"
	if urlToSave == "" {
		return fmt.Errorf("op: %s - empty url to save", op)
	}

	if alias == "" {
		alias = getUniqueAlias(urlToSave)
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

func (u *urlSaver) GetURL(alias string) (*URL, error) {
	const op = "service.shortener.urlSaver.GetURL"

	storeURL, err := u.storage.GetURL(alias)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return converter.StorageToService(storeURL), nil
}
