package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rangzen/gotth-top/services"
	"github.com/rangzen/gotth-top/views/general"
)

type GeneralService interface {
	GeneralStat() (services.GeneralStat, error)
}

type GeneralHandler struct {
	generalService GeneralService
}

func NewGeneralHandler(gs GeneralService) *GeneralHandler {
	return &GeneralHandler{
		generalService: gs,
	}
}

func (gh *GeneralHandler) Base(ctx echo.Context) error {
	gs, err := gh.generalService.GeneralStat()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return render(ctx, general.Layout(gs))
}
