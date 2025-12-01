package urlhandler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rezaabaskhanian/go-url-shortener/internal/param"
)

func (h Handler) CreateUrl(c echo.Context) error {
	var req param.UrlRequest

	// 1️⃣ Bind از body
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
	}

	// // 2️⃣ Parse duration
	// timeDuration, err := time.ParseDuration(req.ExpireAt)
	// if err != nil {
	// 	return echo.NewHTTPError(http.StatusBadRequest, "invalid duration format")
	// }

	// expireAt := time.Now().Add(timeDuration)

	// req.ExpireAt = expireAt.String()

	// 3️⃣ Call Service
	res, err := h.urlSvc.CreateUrl(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("cannot create url: %v", err))
	}

	// 4️⃣ Return JSON
	return c.JSON(http.StatusCreated, res)

}
