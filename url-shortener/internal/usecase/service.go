package usecase

import (
	"fmt"
	"time"

	"github.com/rezaabaskhanian/go-url-shortener/internal/entity"
	"github.com/rezaabaskhanian/go-url-shortener/internal/param"
)

type Repository interface {
	Create(original string, shortCode string, expireAt time.Time) (int64, error)
	GetByShortCode(shortCode string) (entity.URL, error)
	ShowAll() ([]entity.URL, error)
}

type Service struct {
	repo Repository
}

func New(repo Repository) Service {
	return Service{repo: repo}
}

func (s Service) ShowAllUser() ([]entity.URL, error) {

	res, err := s.repo.ShowAll()

	if err != nil {
		return []entity.URL{}, err
	}

	return res, err

}

func (s Service) CreateUrl(req param.UrlRequest) (entity.URL, error) {
	// 1️⃣ Parse duration به expireAt (optional)

	var expireAt time.Time
	if req.ExpireAt != "" {
		duration, err := time.ParseDuration(req.ExpireAt)
		if err != nil {
			return entity.URL{}, fmt.Errorf("invalid expire duration: %w", err)
		}
		t := time.Now().Add(duration)
		expireAt = t
	}

	// 2️⃣ Call repository
	id, err := s.repo.Create(req.Original, req.ShortCode, expireAt)
	if err != nil {
		return entity.URL{}, err
	}

	// 3️⃣ Return entity.URL
	url := entity.URL{
		ID:        id,
		Original:  req.Original,
		ShortCode: req.ShortCode,
		CreatedAt: time.Now(),
		ExpireAt:  expireAt,
	}

	return url, nil
}

func (s Service) GetByShortCode(req param.ShortCodeRequst) (entity.URL, error) {

	res, err := s.repo.GetByShortCode(req.ShortCode)
	if err != nil {
		return entity.URL{}, err
	}

	return res, nil
}
