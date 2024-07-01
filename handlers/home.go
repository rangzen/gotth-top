package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/rangzen/gotth-top/views/home"
)

func Home(ctx echo.Context) error {
	return render(ctx, home.Home())
}
