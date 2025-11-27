package matchingservice

import (
	"game_app/internal/entity"
	"game_app/internal/param"
	"time"
)

type Config struct {
	WaitingTimeout time.Duration `koanf:"waiting_timeout"`
}

type Repo interface {
	AddToWaitingList(userID uint, category entity.Category) error
}

type Service struct {
	repo Repo
}

func New() Service {
	return Service{}
}

func (s Service) AddToWaitingList(req param.AddToWaitingListRequest) (param.AddToWaitingListResponse, error) {

	err := s.repo.AddToWaitingList(req.UserID, req.Category)

	if err != nil {
		return param.AddToWaitingListResponse{}, err
	}

	return param.AddToWaitingListResponse{}, nil

}
