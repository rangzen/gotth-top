package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rangzen/gotth-top/services"
	"github.com/rangzen/gotth-top/views/mem"
)

type MemService interface {
	VirtualMemoryStat() (services.VirtualMemoryStat, error)
}

type MemHandler struct {
	memService MemService
}

func NewMemHandler(ms MemService) *MemHandler {
	return &MemHandler{
		memService: ms,
	}
}

func (mh *MemHandler) ListMem(ctx echo.Context) error {
	vm, err := mh.memService.VirtualMemoryStat()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return render(ctx, mem.Mem(vm))
}
