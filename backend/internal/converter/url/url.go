package url

import (
	"github.com/goriiin/myapp/backend/internal/repository/postgres"
	service "github.com/goriiin/myapp/backend/internal/service "
)

func StorageToService(url *postgres.URL) *service.URL {
	return &service.URL{
		Url:   url.Url,
		Alias: url.Alias,
	}
}
