package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/rangzen/gotha-top/views/home"
)

func Home(ctx echo.Context) error {
	return render(ctx, home.Home())
}
