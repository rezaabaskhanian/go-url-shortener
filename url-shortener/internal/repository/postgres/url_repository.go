package postgres

import (
	"time"

	"github.com/rezaabaskhanian/go-url-shortener/internal/entity"
)

type UrlRepository interface {
	Create(original string, shortCode string, expireAt time.Time) (int64, error)
	GetByShortCode(shortCode string) (entity.URL, error)
	ShowAll() ([]entity.URL, error)
	DeleteExpiredURLs() ([]entity.URL, error)
}
