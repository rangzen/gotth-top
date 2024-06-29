package handlers

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/rangzen/gotha-top/services"
	"github.com/rangzen/gotha-top/views/cpu"
)

type CpuService interface {
	CpuInfoStat(interval time.Duration) (services.CpuInfoStat, error)
}

type CpuHandler struct {
	cpuService CpuService
}

func NewCpuHandler(cs CpuService) *CpuHandler {
	return &CpuHandler{
		cpuService: cs,
	}
}

func (ch *CpuHandler) ListCpuPercents(ctx echo.Context) error {
	cis, err := ch.cpuService.CpuInfoStat(1 * time.Second)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return render(ctx, cpu.Percents(cis.Percents))
}
