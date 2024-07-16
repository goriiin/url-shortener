package service_

// URLSaver TODO: ха-ха надо бы сигнатуру повторить
type URLSaver interface {
	GetURL(alias string) (*Url, error)
	SaveURL(urlToSave string, alias string) error
	RemoveURL(urlToRemove string) error
	EditURL(savedURL string, newAlias string) error
}
