package service_

type URLService interface {
	GetURL(alias string) (*url, error)
	SaveURL(urlToSave string, alias string) error
	RemoveURL(urlToRemove string) error
	EditURL(savedURL string, newAlias string) error
}
