package redis

import "github.com/rezaabaskhanian/go-url-shortener/internal/entity"

type UrlCache interface {
	Set(url entity.URL) error
	Get(shortCode string) (entity.URL, error)
}
