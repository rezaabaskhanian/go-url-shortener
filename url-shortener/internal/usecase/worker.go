package usecase

import (
	"context"
	"log"
	"time"
)

func (s *Service) StartCleanUp(ctx context.Context, interval time.Duration) {

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			s.cleanupUpExpired()
			// when main finish ctx is done or ctx get timeout
		case <-ctx.Done():
			log.Println("Cleanup goroutine stopped")
			return
		}
	}

}

func (s *Service) cleanupUpExpired() {
	expiredUrls, err := s.repo.DeleteExpiredURLs()
	if err != nil {
		log.Println("Error deleting expired URLs from Postgres:", err)
		return
	}

	for _, u := range expiredUrls {
		if err := s.redisRepo.DELETE(u.ShortCode); err != nil {
			log.Println("Error deleting URL from Redis:", u.ShortCode, err)
		}
	}
	log.Printf("Cleaned %d expired URLs\n", len(expiredUrls))

}
