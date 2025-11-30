package entity

import "time"

type URL struct {
	ID        int64
	Original  string
	ShortCode string
	CreatedAt time.Time
	ExpireAt  time.Time
}
