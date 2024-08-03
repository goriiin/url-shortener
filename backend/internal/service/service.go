package service

type URLSaver interface {
	GetURL(alias string) (*Url, error)
	SaveURL(urlToSave string, alias string) (*string, error)
	RemoveURL(urlToRemove string) error
	EditURL(savedURL string, newAlias string) (*string, error)
}
