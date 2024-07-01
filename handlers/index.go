package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/rangzen/gotth-top/views/layout"
)

func Index(ctx echo.Context) error {
	return render(ctx, layout.Base())
}
