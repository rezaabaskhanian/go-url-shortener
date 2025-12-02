package redis

import (
	"github.com/redis/go-redis/v9"
	"github.com/rezaabaskhanian/go-url-shortener/internal/config"
)

func NewRedis(cfg config.RedisConfig) *redis.Client {

	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,     // مثال: "localhost:6379"
		Password: cfg.Password, // بدون رمز: ""
		DB:       cfg.DB,       // معمولاً 0
	})

	return rdb
}
