package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rangzen/gotth-top/services"
	"github.com/rangzen/gotth-top/views/proc"
)

type ProcService interface {
	Processes() ([]services.ProcStat, error)
}

type ProcHandler struct {
	procService ProcService
}

func NewProcHandler(ps ProcService) *ProcHandler {
	return &ProcHandler{
		procService: ps,
	}
}

func (ph *ProcHandler) ListProc(ctx echo.Context) error {
	procs, err := ph.procService.Processes()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return render(ctx, proc.Procs(procs))
}
