package services

import (
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
)

func NewCpuService() *CpuService {
	return &CpuService{}
}

type CpuService struct {
}

type CpuInfoStat struct {
	Interval time.Duration
	Percents []float64
}

func (c *CpuService) CpuInfoStat(interval time.Duration) (CpuInfoStat, error) {
	percent, err := cpu.Percent(interval, true)
	if err != nil {
		return CpuInfoStat{}, err
	}

	cis := CpuInfoStat{
		Interval: interval,
		Percents: percent,
	}

	return cis, nil
}
