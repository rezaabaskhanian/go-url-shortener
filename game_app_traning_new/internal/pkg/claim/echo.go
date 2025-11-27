package claim

import (
	"game_app/internal/pkg/constant"
	"game_app/internal/service/authservice"

	"github.com/labstack/echo/v4"
)

func GetClaims(c echo.Context) *authservice.Claims {

	return c.Get(constant.AuthMiddlewareContextKey).(*authservice.Claims)

}
