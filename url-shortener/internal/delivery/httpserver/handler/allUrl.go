package urlhandler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h Handler) GetAllUrl(c echo.Context) error {

	res, err := h.urlSvc.ShowAllUser()

	if err != nil {
		log.Fatalf("dont get error from databse %v", err)
	}

	return c.JSON(http.StatusOK, res)

}
