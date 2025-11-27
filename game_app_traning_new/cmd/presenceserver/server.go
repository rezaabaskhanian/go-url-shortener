package main

import (
	"game_app/internal/adapter/redis"
	"game_app/internal/delivery/grpcserver/presenceserver"
	"game_app/internal/repository/redis/redispresence"
	"game_app/internal/service/presenceservice"
	"time"
)

func main() {

	redisConfig := redis.Config{
		Host:     "localhost",
		Port:     6379,
		Password: "",
		DB:       0,
	}

	presenceConfig := presenceservice.Config{
		PresenceExpirationTime: 30 * time.Minute,
		PresencePrefix:         "presence",
	}

	redisAdapter := redis.New(redisConfig)
	presenceRepo := redispresence.New(redisAdapter)
	presenceSvc := presenceservice.New(presenceRepo, presenceConfig)

	server := presenceserver.New(presenceSvc)

	server.Start()
}
