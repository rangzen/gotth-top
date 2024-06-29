package services

import (
	"github.com/shirou/gopsutil/v4/mem"
)

func NewMemService() *MemService {
	return &MemService{}
}

type MemService struct {
}

type VirtualMemoryStat struct {
	Total       uint64
	Used        uint64
	UsedPercent float64
}

func (c *MemService) VirtualMemoryStat() (VirtualMemoryStat, error) {
	v, err := mem.VirtualMemory()
	if err != nil {
		return VirtualMemoryStat{}, err
	}

	vms := VirtualMemoryStat{
		Total:       v.Total,
		Used:        v.Used,
		UsedPercent: v.UsedPercent,
	}

	return vms, nil
}
