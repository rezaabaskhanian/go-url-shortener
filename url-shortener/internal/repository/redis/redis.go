package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/rezaabaskhanian/go-url-shortener/internal/entity"
)

type UrlMyRedisRepo struct {
	DB *redis.Client
}

func NeWMyRedisClient(db *redis.Client) UrlCache {
	return UrlMyRedisRepo{DB: db}
}

// Get implements UrlCache.
func (r UrlMyRedisRepo) Get(shortCode string) (entity.URL, error) {
	var url entity.URL

	val, err := r.DB.Get(context.Background(), shortCode).Result()

	if err != nil {
		if err == redis.Nil {
			// وقتی داده در Redis موجود نیست
			return entity.URL{}, fmt.Errorf("not found in cache: %s", shortCode)
		}
		return entity.URL{}, err
	}

	if err := json.Unmarshal([]byte(val), &url); err != nil {
		return entity.URL{}, fmt.Errorf("unmarshal error: %w", err)
	}

	return url, nil
}

// Set implements UrlCache.
func (r UrlMyRedisRepo) Set(url entity.URL) error {

	key := url.ShortCode

	data, err := json.Marshal(url)
	if err != nil {
		return err
	}

	expiration := time.Until(url.ExpireAt)

	err = r.DB.Set(context.Background(), key, data, expiration)
	if err != nil {
		return err
	}

	return nil

}
