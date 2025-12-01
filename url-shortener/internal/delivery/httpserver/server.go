package httpserver

import (
	"github.com/labstack/echo/v4"
	urlhandler "github.com/rezaabaskhanian/go-url-shortener/internal/delivery/httpserver/handler"
)

type Server struct {
	handler urlhandler.Handler
}

func New(handler urlhandler.Handler) Server {
	return Server{handler: handler}
}

func (s Server) Serve() {

	e := echo.New()

	e.GET("/AllUrl", s.handler.GetAllUrl)
	e.POST("/CreateUrl", s.handler.CreateUrl)
	e.POST("/GetShortCode", s.handler.GetByShrotCode)

	e.Logger.Fatal(e.Start(":8082"))

}
