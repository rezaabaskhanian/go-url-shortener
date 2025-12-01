package urlhandler

import "github.com/rezaabaskhanian/go-url-shortener/internal/usecase"

type Handler struct {
	urlSvc usecase.Service
}

func New(urlSvc usecase.Service) Handler {
	return Handler{
		urlSvc: urlSvc,
	}
}
