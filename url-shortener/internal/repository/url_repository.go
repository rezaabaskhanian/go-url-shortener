package repository

import (
	"time"

	"github.com/rezaabaskhanian/go-url-shortener/internal/entity"
)

type UrlRepository interface {
	Create(original string, shortCode string, expireAt time.Time) (int64, error)
	GetByShortCode(shortCode string) (string, error)
	ShowAll() ([]entity.URL, error)
}
