package urlhandler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rezaabaskhanian/go-url-shortener/internal/param"
)

func (h Handler) GetByShrotCode(c echo.Context) error {

	var req param.ShortCodeRequst

	// bind body JSON
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
	}

	res, err := h.urlSvc.GetByShortCode(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("dont show shortcode: %v", err))
	}

	return c.JSON(http.StatusOK, res)
}
